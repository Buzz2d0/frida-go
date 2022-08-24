package fridago

//#include "frida-core.h"
import "C"

type Script struct {
	ptr *C.FridaScript

	Name string
}

func NewScript(s *C.FridaScript, name string) *Script {
	return &Script{
		ptr:  s,
		Name: name,
	}
}

func (s *Script) Free() {
	C.g_object_unref(C.gpointer(s.ptr))
	s.ptr = nil
}

func (s *Script) Load() error {
	var gerr *C.GError
	C.frida_script_load_sync(s.ptr, nil, &gerr)
	if gerr != nil {
		return NewGError(gerr)
	}
	return nil
}

func (s *Script) UnLoad() error {
	var gerr *C.GError
	C.frida_script_unload_sync(s.ptr, nil, &gerr)
	if gerr != nil {
		return NewGError(gerr)
	}
	s.Free()
	return nil
}
