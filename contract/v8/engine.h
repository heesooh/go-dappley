#include <stdint.h>
#include <stdbool.h>
#include "lib/transaction_struct.h"
#include "lib/utxo_struct.h"
#include "lib/vm_error.h"

#ifdef WIN32 
#ifdef V8DLL
#define EXPORT __declspec(dllexport)
#else 
#define EXPORT __declspec(dllimport)
#endif
#else
#define EXPORT __attribute__((__visibility__("default")))
#endif 

#define BUILD_MATH            0x0000000000000001
#define BUILD_MATH_RANDOM       0x0000000000000002
#define BUILD_BLOCKCHAIN        0x0000000000000004
#define BUILD_DEFAULT_VER (BUILD_MATH | BUILD_BLOCKCHAIN)

#ifndef _DAPPLEY_NF_VM_V8_ENGINE_H_
#define _DAPPLEY_NF_VM_V8_ENGINE_H_
typedef struct V8EngineStats {
  size_t count_of_executed_instructions;
  size_t total_memory_size;
  size_t total_heap_size;
  size_t total_heap_size_executable;
  size_t total_physical_size;
  size_t total_available_size;
  size_t used_heap_size;
  size_t heap_size_limit;
  size_t malloced_memory;
  size_t peak_malloced_memory;
  size_t total_array_buffer_size;
  size_t peak_array_buffer_size;
} V8EngineStats;

typedef struct V8Engine {

  void *isolate;
  void *allocator;
  size_t limits_of_executed_instructions;
  size_t limits_of_total_memory_size;
  bool is_requested_terminate_execution;
  bool is_unexpected_error_happen;
  bool is_inner_nvm_error_happen;
  int testing;
  int timeout;
  uint64_t ver;
  V8EngineStats stats;

} V8Engine;
#endif

#ifdef __cplusplus
extern "C" {
#endif
    typedef bool (*FuncVerifyAddress)(const char *address);
    typedef int (*FuncTransfer)(void *handler, const char *to, const char *amount, const char *tip);
    typedef char* (*FuncStorageGet)(void *address, const char *key);
    typedef int (*FuncStorageSet)(void *address, const char *key, const char *value);
    typedef int (*FuncStorageDel)(void *address, const char *key);
    typedef int (*FuncTriggerEvent)(void *address, const char *topic, const char *data);
    typedef void (*FuncTransactionGet)(void* address, void* context);
    typedef void (*FuncPrevUtxoGet)(void* address, void* context);
    typedef void (*FuncLogger)(unsigned int level, char** args, int length);
    typedef int (*FuncRecordReward)(void *handler, const char *address, const char *amount);
    typedef bool (*FuncVerifySignature)(const char *msg, const char *pubKey, const char *sig);
    typedef bool (*FuncVerifyPublicKey)(const char *addr, const char *pubKey);
    typedef int (*FuncRandom)(void *handler, int max);
    typedef int (*FuncGetCurrBlockHeight)(void *handler);
    typedef char* (*FuncGetNodeAddress)(void *handler);
	typedef void* (*FuncMalloc)(size_t size);
	typedef void  (*FuncFree)(void* data);

    EXPORT V8Engine *CreateEngine();
    EXPORT void Initialize();
    EXPORT int executeV8Script(const char *sourceCode, uintptr_t handler, char **result, V8Engine *e);
    EXPORT void InitializeBlockchain(FuncVerifyAddress verifyAddress, FuncTransfer transfer, FuncGetCurrBlockHeight getCurrBlockHeight, FuncGetNodeAddress getNodeAddress);
    EXPORT void InitializeRewardDistributor(FuncRecordReward recordReward);
    EXPORT void InitializeStorage(FuncStorageGet get, FuncStorageSet set, FuncStorageDel del);
    EXPORT void InitializeEvent(FuncTriggerEvent triggerEvent);
    EXPORT void InitializeTransaction(FuncTransactionGet get);
    EXPORT void InitializeCrypto(FuncVerifySignature verifySignature, FuncVerifyPublicKey verifyPublicKey);
    EXPORT void InitializeMath(FuncRandom random);
    EXPORT void SetTransactionData(struct transaction_t* tx, void* context);
    EXPORT void InitializePrevUtxo(FuncPrevUtxoGet get);
    EXPORT void SetPrevUtxoData(struct utxo_t* utxos, int length, void* context);
    EXPORT void InitializeLogger(FuncLogger logger);
    EXPORT void InitializeSmartContract(char* source);
    EXPORT void DisposeV8();
	EXPORT void InitializeMemoryFunc(FuncMalloc mallocFunc, FuncFree freeFunc);
	EXPORT void DeleteEngine(V8Engine *e);

#ifdef __cplusplus
}
#endif

