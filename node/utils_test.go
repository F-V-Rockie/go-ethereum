// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

// Contains a batch of utility type declarations used by the tests. As the node
// operates on unique types, a lot of them are needed to check various features.

package node

import (
	"reflect"

	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/rpc"
)

// NoopService is a trivial implementation of the Service interface.
type NoopService struct{}

func (s *NoopService) Protocols() []p2p.Protocol { log.DebugLog() return nil }
func (s *NoopService) APIs() []rpc.API           { log.DebugLog() return nil }
func (s *NoopService) Start(*p2p.Server) error   { log.DebugLog() return nil }
func (s *NoopService) Stop() error               { log.DebugLog() return nil }

func NewNoopService(*ServiceContext) (Service, error) { log.DebugLog() return new(NoopService), nil }

// Set of services all wrapping the base NoopService resulting in the same method
// signatures but different outer types.
type NoopServiceA struct{ NoopService }
type NoopServiceB struct{ NoopService }
type NoopServiceC struct{ NoopService }

func NewNoopServiceA(*ServiceContext) (Service, error) { log.DebugLog() return new(NoopServiceA), nil }
func NewNoopServiceB(*ServiceContext) (Service, error) { log.DebugLog() return new(NoopServiceB), nil }
func NewNoopServiceC(*ServiceContext) (Service, error) { log.DebugLog() return new(NoopServiceC), nil }

// InstrumentedService is an implementation of Service for which all interface
// methods can be instrumented both return value as well as event hook wise.
type InstrumentedService struct {
	protocols []p2p.Protocol
	apis      []rpc.API
	start     error
	stop      error

	protocolsHook func()
	startHook     func(*p2p.Server)
	stopHook      func()
}

func NewInstrumentedService(*ServiceContext) (Service, error) { log.DebugLog() return new(InstrumentedService), nil }

func (s *InstrumentedService) Protocols() []p2p.Protocol { log.DebugLog()
	if s.protocolsHook != nil {
		s.protocolsHook()
	}
	return s.protocols
}

func (s *InstrumentedService) APIs() []rpc.API { log.DebugLog()
	return s.apis
}

func (s *InstrumentedService) Start(server *p2p.Server) error { log.DebugLog()
	if s.startHook != nil {
		s.startHook(server)
	}
	return s.start
}

func (s *InstrumentedService) Stop() error { log.DebugLog()
	if s.stopHook != nil {
		s.stopHook()
	}
	return s.stop
}

// InstrumentingWrapper is a method to specialize a service constructor returning
// a generic InstrumentedService into one returning a wrapping specific one.
type InstrumentingWrapper func(base ServiceConstructor) ServiceConstructor

func InstrumentingWrapperMaker(base ServiceConstructor, kind reflect.Type) ServiceConstructor { log.DebugLog()
	return func(ctx *ServiceContext) (Service, error) {
		obj, err := base(ctx)
		if err != nil {
			return nil, err
		}
		wrapper := reflect.New(kind)
		wrapper.Elem().Field(0).Set(reflect.ValueOf(obj).Elem())

		return wrapper.Interface().(Service), nil
	}
}

// Set of services all wrapping the base InstrumentedService resulting in the
// same method signatures but different outer types.
type InstrumentedServiceA struct{ InstrumentedService }
type InstrumentedServiceB struct{ InstrumentedService }
type InstrumentedServiceC struct{ InstrumentedService }

func InstrumentedServiceMakerA(base ServiceConstructor) ServiceConstructor { log.DebugLog()
	return InstrumentingWrapperMaker(base, reflect.TypeOf(InstrumentedServiceA{}))
}

func InstrumentedServiceMakerB(base ServiceConstructor) ServiceConstructor { log.DebugLog()
	return InstrumentingWrapperMaker(base, reflect.TypeOf(InstrumentedServiceB{}))
}

func InstrumentedServiceMakerC(base ServiceConstructor) ServiceConstructor { log.DebugLog()
	return InstrumentingWrapperMaker(base, reflect.TypeOf(InstrumentedServiceC{}))
}

// OneMethodApi is a single-method API handler to be returned by test services.
type OneMethodApi struct {
	fun func()
}

func (api *OneMethodApi) TheOneMethod() { log.DebugLog()
	if api.fun != nil {
		api.fun()
	}
}
