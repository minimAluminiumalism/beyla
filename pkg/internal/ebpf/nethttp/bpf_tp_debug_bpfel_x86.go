// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64
// +build 386 amd64

package nethttp

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type bpf_tp_debugGoroutineMetadata struct {
	Parent    uint64
	Timestamp uint64
}

type bpf_tp_debugHttpFuncInvocationT struct {
	StartMonotimeNs uint64
	ReqPtr          uint64
	Tp              bpf_tp_debugTpInfoT
}

type bpf_tp_debugPidKeyT struct {
	Pid       uint32
	Namespace uint32
}

type bpf_tp_debugTpInfoT struct {
	TraceId  [16]uint8
	SpanId   [8]uint8
	ParentId [8]uint8
	Ts       uint64
	Flags    uint8
	_        [7]byte
}

// loadBpf_tp_debug returns the embedded CollectionSpec for bpf_tp_debug.
func loadBpf_tp_debug() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_Bpf_tp_debugBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load bpf_tp_debug: %w", err)
	}

	return spec, err
}

// loadBpf_tp_debugObjects loads bpf_tp_debug and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*bpf_tp_debugObjects
//	*bpf_tp_debugPrograms
//	*bpf_tp_debugMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadBpf_tp_debugObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadBpf_tp_debug()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// bpf_tp_debugSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpf_tp_debugSpecs struct {
	bpf_tp_debugProgramSpecs
	bpf_tp_debugMapSpecs
}

// bpf_tp_debugSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpf_tp_debugProgramSpecs struct {
	UprobeServeHTTP          *ebpf.ProgramSpec `ebpf:"uprobe_ServeHTTP"`
	UprobeWriteHeader        *ebpf.ProgramSpec `ebpf:"uprobe_WriteHeader"`
	UprobeReadRequestReturns *ebpf.ProgramSpec `ebpf:"uprobe_readRequestReturns"`
	UprobeRoundTrip          *ebpf.ProgramSpec `ebpf:"uprobe_roundTrip"`
	UprobeRoundTripReturn    *ebpf.ProgramSpec `ebpf:"uprobe_roundTripReturn"`
	UprobeWriteSubset        *ebpf.ProgramSpec `ebpf:"uprobe_writeSubset"`
}

// bpf_tp_debugMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpf_tp_debugMapSpecs struct {
	Events                    *ebpf.MapSpec `ebpf:"events"`
	GoTraceMap                *ebpf.MapSpec `ebpf:"go_trace_map"`
	GolangMapbucketStorageMap *ebpf.MapSpec `ebpf:"golang_mapbucket_storage_map"`
	HeaderReqMap              *ebpf.MapSpec `ebpf:"header_req_map"`
	OngoingGoroutines         *ebpf.MapSpec `ebpf:"ongoing_goroutines"`
	OngoingHttpClientRequests *ebpf.MapSpec `ebpf:"ongoing_http_client_requests"`
	OngoingHttpServerRequests *ebpf.MapSpec `ebpf:"ongoing_http_server_requests"`
	PidCache                  *ebpf.MapSpec `ebpf:"pid_cache"`
	ValidPids                 *ebpf.MapSpec `ebpf:"valid_pids"`
}

// bpf_tp_debugObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadBpf_tp_debugObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpf_tp_debugObjects struct {
	bpf_tp_debugPrograms
	bpf_tp_debugMaps
}

func (o *bpf_tp_debugObjects) Close() error {
	return _Bpf_tp_debugClose(
		&o.bpf_tp_debugPrograms,
		&o.bpf_tp_debugMaps,
	)
}

// bpf_tp_debugMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadBpf_tp_debugObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpf_tp_debugMaps struct {
	Events                    *ebpf.Map `ebpf:"events"`
	GoTraceMap                *ebpf.Map `ebpf:"go_trace_map"`
	GolangMapbucketStorageMap *ebpf.Map `ebpf:"golang_mapbucket_storage_map"`
	HeaderReqMap              *ebpf.Map `ebpf:"header_req_map"`
	OngoingGoroutines         *ebpf.Map `ebpf:"ongoing_goroutines"`
	OngoingHttpClientRequests *ebpf.Map `ebpf:"ongoing_http_client_requests"`
	OngoingHttpServerRequests *ebpf.Map `ebpf:"ongoing_http_server_requests"`
	PidCache                  *ebpf.Map `ebpf:"pid_cache"`
	ValidPids                 *ebpf.Map `ebpf:"valid_pids"`
}

func (m *bpf_tp_debugMaps) Close() error {
	return _Bpf_tp_debugClose(
		m.Events,
		m.GoTraceMap,
		m.GolangMapbucketStorageMap,
		m.HeaderReqMap,
		m.OngoingGoroutines,
		m.OngoingHttpClientRequests,
		m.OngoingHttpServerRequests,
		m.PidCache,
		m.ValidPids,
	)
}

// bpf_tp_debugPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadBpf_tp_debugObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpf_tp_debugPrograms struct {
	UprobeServeHTTP          *ebpf.Program `ebpf:"uprobe_ServeHTTP"`
	UprobeWriteHeader        *ebpf.Program `ebpf:"uprobe_WriteHeader"`
	UprobeReadRequestReturns *ebpf.Program `ebpf:"uprobe_readRequestReturns"`
	UprobeRoundTrip          *ebpf.Program `ebpf:"uprobe_roundTrip"`
	UprobeRoundTripReturn    *ebpf.Program `ebpf:"uprobe_roundTripReturn"`
	UprobeWriteSubset        *ebpf.Program `ebpf:"uprobe_writeSubset"`
}

func (p *bpf_tp_debugPrograms) Close() error {
	return _Bpf_tp_debugClose(
		p.UprobeServeHTTP,
		p.UprobeWriteHeader,
		p.UprobeReadRequestReturns,
		p.UprobeRoundTrip,
		p.UprobeRoundTripReturn,
		p.UprobeWriteSubset,
	)
}

func _Bpf_tp_debugClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed bpf_tp_debug_bpfel_x86.o
var _Bpf_tp_debugBytes []byte
