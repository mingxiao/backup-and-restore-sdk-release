// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	gcs "github.com/cloudfoundry-incubator/gcs-blobstore-backup-restore"
)

type FakeBucket struct {
	NameStub        func() string
	nameMutex       sync.RWMutex
	nameArgsForCall []struct{}
	nameReturns     struct {
		result1 string
	}
	nameReturnsOnCall map[int]struct {
		result1 string
	}
	ListBlobsStub        func() ([]gcs.Blob, error)
	listBlobsMutex       sync.RWMutex
	listBlobsArgsForCall []struct{}
	listBlobsReturns     struct {
		result1 []gcs.Blob
		result2 error
	}
	listBlobsReturnsOnCall map[int]struct {
		result1 []gcs.Blob
		result2 error
	}
	LastBackupBlobsStub        func() (map[string]gcs.Blob, error)
	lastBackupBlobsMutex       sync.RWMutex
	lastBackupBlobsArgsForCall []struct{}
	lastBackupBlobsReturns     struct {
		result1 map[string]gcs.Blob
		result2 error
	}
	lastBackupBlobsReturnsOnCall map[int]struct {
		result1 map[string]gcs.Blob
		result2 error
	}
	CopyBlobWithinBucketStub        func(string, string) error
	copyBlobWithinBucketMutex       sync.RWMutex
	copyBlobWithinBucketArgsForCall []struct {
		arg1 string
		arg2 string
	}
	copyBlobWithinBucketReturns struct {
		result1 error
	}
	copyBlobWithinBucketReturnsOnCall map[int]struct {
		result1 error
	}
	CopyBlobBetweenBucketsStub        func(gcs.Bucket, string, string) error
	copyBlobBetweenBucketsMutex       sync.RWMutex
	copyBlobBetweenBucketsArgsForCall []struct {
		arg1 gcs.Bucket
		arg2 string
		arg3 string
	}
	copyBlobBetweenBucketsReturns struct {
		result1 error
	}
	copyBlobBetweenBucketsReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteBlobStub        func(string) error
	deleteBlobMutex       sync.RWMutex
	deleteBlobArgsForCall []struct {
		arg1 string
	}
	deleteBlobReturns struct {
		result1 error
	}
	deleteBlobReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBucket) Name() string {
	fake.nameMutex.Lock()
	ret, specificReturn := fake.nameReturnsOnCall[len(fake.nameArgsForCall)]
	fake.nameArgsForCall = append(fake.nameArgsForCall, struct{}{})
	fake.recordInvocation("Name", []interface{}{})
	fake.nameMutex.Unlock()
	if fake.NameStub != nil {
		return fake.NameStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.nameReturns.result1
}

func (fake *FakeBucket) NameCallCount() int {
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	return len(fake.nameArgsForCall)
}

func (fake *FakeBucket) NameReturns(result1 string) {
	fake.NameStub = nil
	fake.nameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeBucket) NameReturnsOnCall(i int, result1 string) {
	fake.NameStub = nil
	if fake.nameReturnsOnCall == nil {
		fake.nameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.nameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeBucket) ListBlobs() ([]gcs.Blob, error) {
	fake.listBlobsMutex.Lock()
	ret, specificReturn := fake.listBlobsReturnsOnCall[len(fake.listBlobsArgsForCall)]
	fake.listBlobsArgsForCall = append(fake.listBlobsArgsForCall, struct{}{})
	fake.recordInvocation("ListBlobs", []interface{}{})
	fake.listBlobsMutex.Unlock()
	if fake.ListBlobsStub != nil {
		return fake.ListBlobsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listBlobsReturns.result1, fake.listBlobsReturns.result2
}

func (fake *FakeBucket) ListBlobsCallCount() int {
	fake.listBlobsMutex.RLock()
	defer fake.listBlobsMutex.RUnlock()
	return len(fake.listBlobsArgsForCall)
}

func (fake *FakeBucket) ListBlobsReturns(result1 []gcs.Blob, result2 error) {
	fake.ListBlobsStub = nil
	fake.listBlobsReturns = struct {
		result1 []gcs.Blob
		result2 error
	}{result1, result2}
}

func (fake *FakeBucket) ListBlobsReturnsOnCall(i int, result1 []gcs.Blob, result2 error) {
	fake.ListBlobsStub = nil
	if fake.listBlobsReturnsOnCall == nil {
		fake.listBlobsReturnsOnCall = make(map[int]struct {
			result1 []gcs.Blob
			result2 error
		})
	}
	fake.listBlobsReturnsOnCall[i] = struct {
		result1 []gcs.Blob
		result2 error
	}{result1, result2}
}

func (fake *FakeBucket) LastBackupBlobs() (map[string]gcs.Blob, error) {
	fake.lastBackupBlobsMutex.Lock()
	ret, specificReturn := fake.lastBackupBlobsReturnsOnCall[len(fake.lastBackupBlobsArgsForCall)]
	fake.lastBackupBlobsArgsForCall = append(fake.lastBackupBlobsArgsForCall, struct{}{})
	fake.recordInvocation("LastBackupBlobs", []interface{}{})
	fake.lastBackupBlobsMutex.Unlock()
	if fake.LastBackupBlobsStub != nil {
		return fake.LastBackupBlobsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.lastBackupBlobsReturns.result1, fake.lastBackupBlobsReturns.result2
}

func (fake *FakeBucket) LastBackupBlobsCallCount() int {
	fake.lastBackupBlobsMutex.RLock()
	defer fake.lastBackupBlobsMutex.RUnlock()
	return len(fake.lastBackupBlobsArgsForCall)
}

func (fake *FakeBucket) LastBackupBlobsReturns(result1 map[string]gcs.Blob, result2 error) {
	fake.LastBackupBlobsStub = nil
	fake.lastBackupBlobsReturns = struct {
		result1 map[string]gcs.Blob
		result2 error
	}{result1, result2}
}

func (fake *FakeBucket) LastBackupBlobsReturnsOnCall(i int, result1 map[string]gcs.Blob, result2 error) {
	fake.LastBackupBlobsStub = nil
	if fake.lastBackupBlobsReturnsOnCall == nil {
		fake.lastBackupBlobsReturnsOnCall = make(map[int]struct {
			result1 map[string]gcs.Blob
			result2 error
		})
	}
	fake.lastBackupBlobsReturnsOnCall[i] = struct {
		result1 map[string]gcs.Blob
		result2 error
	}{result1, result2}
}

func (fake *FakeBucket) CopyBlobWithinBucket(arg1 string, arg2 string) error {
	fake.copyBlobWithinBucketMutex.Lock()
	ret, specificReturn := fake.copyBlobWithinBucketReturnsOnCall[len(fake.copyBlobWithinBucketArgsForCall)]
	fake.copyBlobWithinBucketArgsForCall = append(fake.copyBlobWithinBucketArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("CopyBlobWithinBucket", []interface{}{arg1, arg2})
	fake.copyBlobWithinBucketMutex.Unlock()
	if fake.CopyBlobWithinBucketStub != nil {
		return fake.CopyBlobWithinBucketStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.copyBlobWithinBucketReturns.result1
}

func (fake *FakeBucket) CopyBlobWithinBucketCallCount() int {
	fake.copyBlobWithinBucketMutex.RLock()
	defer fake.copyBlobWithinBucketMutex.RUnlock()
	return len(fake.copyBlobWithinBucketArgsForCall)
}

func (fake *FakeBucket) CopyBlobWithinBucketArgsForCall(i int) (string, string) {
	fake.copyBlobWithinBucketMutex.RLock()
	defer fake.copyBlobWithinBucketMutex.RUnlock()
	return fake.copyBlobWithinBucketArgsForCall[i].arg1, fake.copyBlobWithinBucketArgsForCall[i].arg2
}

func (fake *FakeBucket) CopyBlobWithinBucketReturns(result1 error) {
	fake.CopyBlobWithinBucketStub = nil
	fake.copyBlobWithinBucketReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBucket) CopyBlobWithinBucketReturnsOnCall(i int, result1 error) {
	fake.CopyBlobWithinBucketStub = nil
	if fake.copyBlobWithinBucketReturnsOnCall == nil {
		fake.copyBlobWithinBucketReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.copyBlobWithinBucketReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBucket) CopyBlobBetweenBuckets(arg1 gcs.Bucket, arg2 string, arg3 string) error {
	fake.copyBlobBetweenBucketsMutex.Lock()
	ret, specificReturn := fake.copyBlobBetweenBucketsReturnsOnCall[len(fake.copyBlobBetweenBucketsArgsForCall)]
	fake.copyBlobBetweenBucketsArgsForCall = append(fake.copyBlobBetweenBucketsArgsForCall, struct {
		arg1 gcs.Bucket
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("CopyBlobBetweenBuckets", []interface{}{arg1, arg2, arg3})
	fake.copyBlobBetweenBucketsMutex.Unlock()
	if fake.CopyBlobBetweenBucketsStub != nil {
		return fake.CopyBlobBetweenBucketsStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.copyBlobBetweenBucketsReturns.result1
}

func (fake *FakeBucket) CopyBlobBetweenBucketsCallCount() int {
	fake.copyBlobBetweenBucketsMutex.RLock()
	defer fake.copyBlobBetweenBucketsMutex.RUnlock()
	return len(fake.copyBlobBetweenBucketsArgsForCall)
}

func (fake *FakeBucket) CopyBlobBetweenBucketsArgsForCall(i int) (gcs.Bucket, string, string) {
	fake.copyBlobBetweenBucketsMutex.RLock()
	defer fake.copyBlobBetweenBucketsMutex.RUnlock()
	return fake.copyBlobBetweenBucketsArgsForCall[i].arg1, fake.copyBlobBetweenBucketsArgsForCall[i].arg2, fake.copyBlobBetweenBucketsArgsForCall[i].arg3
}

func (fake *FakeBucket) CopyBlobBetweenBucketsReturns(result1 error) {
	fake.CopyBlobBetweenBucketsStub = nil
	fake.copyBlobBetweenBucketsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBucket) CopyBlobBetweenBucketsReturnsOnCall(i int, result1 error) {
	fake.CopyBlobBetweenBucketsStub = nil
	if fake.copyBlobBetweenBucketsReturnsOnCall == nil {
		fake.copyBlobBetweenBucketsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.copyBlobBetweenBucketsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBucket) DeleteBlob(arg1 string) error {
	fake.deleteBlobMutex.Lock()
	ret, specificReturn := fake.deleteBlobReturnsOnCall[len(fake.deleteBlobArgsForCall)]
	fake.deleteBlobArgsForCall = append(fake.deleteBlobArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("DeleteBlob", []interface{}{arg1})
	fake.deleteBlobMutex.Unlock()
	if fake.DeleteBlobStub != nil {
		return fake.DeleteBlobStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteBlobReturns.result1
}

func (fake *FakeBucket) DeleteBlobCallCount() int {
	fake.deleteBlobMutex.RLock()
	defer fake.deleteBlobMutex.RUnlock()
	return len(fake.deleteBlobArgsForCall)
}

func (fake *FakeBucket) DeleteBlobArgsForCall(i int) string {
	fake.deleteBlobMutex.RLock()
	defer fake.deleteBlobMutex.RUnlock()
	return fake.deleteBlobArgsForCall[i].arg1
}

func (fake *FakeBucket) DeleteBlobReturns(result1 error) {
	fake.DeleteBlobStub = nil
	fake.deleteBlobReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBucket) DeleteBlobReturnsOnCall(i int, result1 error) {
	fake.DeleteBlobStub = nil
	if fake.deleteBlobReturnsOnCall == nil {
		fake.deleteBlobReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteBlobReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBucket) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	fake.listBlobsMutex.RLock()
	defer fake.listBlobsMutex.RUnlock()
	fake.lastBackupBlobsMutex.RLock()
	defer fake.lastBackupBlobsMutex.RUnlock()
	fake.copyBlobWithinBucketMutex.RLock()
	defer fake.copyBlobWithinBucketMutex.RUnlock()
	fake.copyBlobBetweenBucketsMutex.RLock()
	defer fake.copyBlobBetweenBucketsMutex.RUnlock()
	fake.deleteBlobMutex.RLock()
	defer fake.deleteBlobMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBucket) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ gcs.Bucket = new(FakeBucket)
