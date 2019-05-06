// Copyright 2015 the V8 project authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <assert.h>
#include <v8.h>
#include <libplatform/libplatform.h>
#include "engine.h"
#include "lib/allocator.h"
#include "lib/instruction_counter.h"
#include "lib/blockchain.h"
#include "lib/load_lib.h"
#include "lib/load_sc.h"
#include "lib/storage.h"
#include "lib/event.h"
#include "lib/logger.h"
#include "lib/transaction.h"
#include "lib/reward_distributor.h"
#include "lib/prev_utxo.h"
#include "lib/crypto.h"
#include "lib/math.h"
#include "lib/memory.h"
#include "lib/vm_error.h"

using namespace v8;
std::unique_ptr<Platform> platformPtr;
static char* wrapReturnResult(const char* src);

#define ExecuteTimeOut  5*1000*1000
void EngineLimitsCheckDelegate(Isolate *isolate, size_t count, void *listenerContext);

void Initialize(){
    // Initialize V8.
    platformPtr = platform::NewDefaultPlatform();
    V8::InitializePlatform(platformPtr.get());
    V8::Initialize();

    // Initialize V8Engine.
    SetInstructionCounterIncrListener(EngineLimitsCheckDelegate);
}

const char* toCString(const v8::String::Utf8Value& value) {
  return *value ? *value : "<string conversion failed>";
}

void reportException(v8::Isolate* isolate, v8::TryCatch* try_catch) {
  v8::HandleScope handle_scope(isolate);
  v8::String::Utf8Value exception(isolate, try_catch->Exception());
  const char* exception_string = toCString(exception);
  v8::Local<v8::Message> message = try_catch->Message();
  if (message.IsEmpty()) {
    // V8 didn't provide any extra information about this error; just
    // print the exception.
    fprintf(stderr, "%s\n", exception_string);
  } else {
    // Print (filename):(line number): (message).
    v8::String::Utf8Value filename(isolate,
      message->GetScriptOrigin().ResourceName());
    v8::Local<v8::Context> context(isolate->GetCurrentContext());
    const char* filename_string = toCString(filename);
    int linenum = message->GetLineNumber(context).FromJust();
    fprintf(stderr, "%s:%i: %s\n", filename_string, linenum, exception_string);
    // Print line of source code.
    v8::String::Utf8Value sourceline(
      isolate, message->GetSourceLine(context).ToLocalChecked());
    const char* sourceline_string = toCString(sourceline);
    fprintf(stderr, "%s\n", sourceline_string);
    // Print wavy underline (GetUnderline is deprecated).
    int start = message->GetStartColumn(context).FromJust();
    for (int i = 0; i < start; i++) {
      fprintf(stderr, " ");
    }
    int end = message->GetEndColumn(context).FromJust();
    for (int i = start; i < end; i++) {
      fprintf(stderr, "^");
    }
    fprintf(stderr, "\n");
    v8::Local<v8::Value> stack_trace_string;
    if (try_catch->StackTrace(context).ToLocal(&stack_trace_string) &&
    stack_trace_string->IsString() &&
    v8::Local<v8::String>::Cast(stack_trace_string)->Length() > 0) {
      v8::String::Utf8Value stack_trace(isolate, stack_trace_string);
      const char* stack_trace_string = toCString(stack_trace);
      fprintf(stderr, "%s\n", stack_trace_string);
    }
  }
}

int executeV8Script(const char *sourceCode, uintptr_t handler, char **result, V8Engine *e) {
  // Create a new Isolate and make it the current one.
  Isolate* isolate = static_cast<Isolate *>(e->isolate);
  Locker locker(isolate);
  int errorCode = 0;

  {
    Isolate::Scope isolate_scope(isolate);

    // Create a stack-allocated handle scope.
    HandleScope handle_scope(isolate);
    //
    Local<ObjectTemplate> globalTpl = NewNativeRequireFunction(isolate);

    // Set up an exception handler
    TryCatch try_catch(isolate);

    // Create a new context.
    Local<Context> context = v8::Context::New(isolate, NULL, globalTpl);

    // Enter the context for compiling and running the hello world script.
    Context::Scope context_scope(context);
    NewBlockchainInstance(isolate, context, (void *)handler);
    NewCryptoInstance(isolate, context, (void *)handler);
    NewStorageInstance(isolate, context, (void *)handler);
    NewLoggerInstance(isolate, context, (void *)handler);
    NewTransactionInstance(isolate, context, (void *)handler);
    NewRewardDistributorInstance(isolate, context, (void *)handler);
    NewPrevUtxoInstance(isolate, context, (void *)handler);
    NewMathInstance(isolate, context, (void *)handler);
    NewEventInstance(isolate, context, (void *)handler);

    NewInstructionCounterInstance(isolate, context,
                                    &(e->stats.count_of_executed_instructions), e);
    LoadLibraries(isolate, context);
    {

      // Create a string containing the JavaScript source code.
      Local<String> source = String::NewFromUtf8(
        isolate,
        sourceCode,
        NewStringType::kNormal
      ).ToLocalChecked();

      // Compile the source code.
      Local<Script> script;
      if (!Script::Compile(context, source).ToLocal(&script)) {
        reportException(isolate, &try_catch);
        *result = wrapReturnResult("1");
        errorCode = 1;
        return errorCode;
      }
      // Run the script to get the result.
      Local<Value> scriptRes;
      if (!script->Run(context).ToLocal(&scriptRes)) {
        assert(try_catch.HasCaught());
        reportException(isolate, &try_catch);
        *result = wrapReturnResult("1");
        errorCode = 1;
        return errorCode;
      }

      // set result.
      if (result != NULL)  {
        Local<Object> obj = scriptRes.As<Object>();
        if (!obj->IsUndefined()) {
          String::Utf8Value str(isolate, obj);
          *result = wrapReturnResult(*str);
        }
      }
    }
  }

  return errorCode;
}

char* wrapReturnResult(const char* src) {
	char* result = (char *)MyMalloc(strlen(src) + 1);
	strcpy(result, src);
	return result;
}

void DisposeV8(){
    V8::Dispose();
    V8::ShutdownPlatform();
    if (platformPtr) {
        platformPtr = NULL;
    }
}

V8Engine *CreateEngine() {
  ArrayBuffer::Allocator *allocator = new ArrayBufferAllocator();

  Isolate::CreateParams create_params;
  create_params.array_buffer_allocator = allocator;

  Isolate *isolate = Isolate::New(create_params);

  V8Engine *e = (V8Engine *)calloc(1, sizeof(V8Engine));
  e->allocator = allocator;
  e->isolate = isolate;
  e->timeout = ExecuteTimeOut;
  e->ver = BUILD_DEFAULT_VER; //default load initial com
  return e;
}

void ReadMemoryStatistics(V8Engine *e) {
  Isolate *isolate = static_cast<Isolate *>(e->isolate);
  ArrayBufferAllocator *allocator =
      static_cast<ArrayBufferAllocator *>(e->allocator);

  HeapStatistics heap_stats;
  isolate->GetHeapStatistics(&heap_stats);

  V8EngineStats *stats = &(e->stats);
  stats->heap_size_limit = heap_stats.heap_size_limit();
  stats->malloced_memory = heap_stats.malloced_memory();
  stats->peak_malloced_memory = heap_stats.peak_malloced_memory();
  stats->total_available_size = heap_stats.total_available_size();
  stats->total_heap_size = heap_stats.total_heap_size();
  stats->total_heap_size_executable = heap_stats.total_heap_size_executable();
  stats->total_physical_size = heap_stats.total_physical_size();
  stats->used_heap_size = heap_stats.used_heap_size();
  stats->total_array_buffer_size = allocator->total_available_size();
  stats->peak_array_buffer_size = allocator->peak_allocated_size();

  stats->total_memory_size =
      stats->total_heap_size + stats->peak_array_buffer_size;
}

void TerminateExecution(V8Engine *e) {
  if (e->is_requested_terminate_execution) {
    return;
  }
  Isolate *isolate = static_cast<Isolate *>(e->isolate);
  isolate->TerminateExecution();
  e->is_requested_terminate_execution = true;
}

void SetInnerContractErrFlag(V8Engine *e) {
  e->is_inner_nvm_error_happen = true;
}

void DeleteEngine(V8Engine *e) {
  Isolate *isolate = static_cast<Isolate *>(e->isolate);
  isolate->Dispose();

  delete static_cast<ArrayBuffer::Allocator *>(e->allocator);

  free(e);
}

int IsEngineLimitsExceeded(V8Engine *e) {
  // TODO: read memory stats everytime may impact the performance.
  ReadMemoryStatistics(e);
  if (e->limits_of_executed_instructions > 0 &&
      e->limits_of_executed_instructions <
          e->stats.count_of_executed_instructions) {
    // Reach instruction limits.
    return VM_GAS_LIMIT_ERR;
  } else if (e->limits_of_total_memory_size > 0 &&
             e->limits_of_total_memory_size < e->stats.total_memory_size) {
    // reach memory limits.
    return VM_MEM_LIMIT_ERR;
  }
  return 0;
}

void EngineLimitsCheckDelegate(Isolate *isolate, size_t count,
                               void *listenerContext) {
  V8Engine *e = static_cast<V8Engine *>(listenerContext);

  if (IsEngineLimitsExceeded(e)) {
    TerminateExecution(e);
  }
}
