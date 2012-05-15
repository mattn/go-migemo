package migemo

/*
#include <migemo.h>
static char* _migemo_query(migemo* object, const char* query) {
  return (char*)migemo_query(object, (const unsigned char*)query);
}
static void _migemo_release(migemo* object, const char* str) {
  migemo_release(object, (unsigned char*)str);
}

*/
// #cgo LDFLAGS: -lmigemo
import "C";

const MIGEMO_DICTID_MIGEMO = 1;

func Open(dict string) *C.migemo {
    return C.migemo_open(C.CString(dict));
}

func Close(m *C.migemo) {
    C.migemo_close(m);
}

func Load(m *C.migemo, dict string) C.int {
    return C.migemo_load(m, MIGEMO_DICTID_MIGEMO, C.CString(dict));
}

func Query(m *C.migemo, q string) string {
    p := C._migemo_query(m, C.CString(q));
	ret := C.GoString(p);
	C._migemo_release(m, p);
	return ret;
}
