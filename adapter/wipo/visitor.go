package wipo

import (
	"github.com/dop251/goja"
)
/****************************************************************
	wipo-visitor-uunid= 设置cookie数据信息
 */
var js = ` //generate unique visitor id cookie
    if (!Math.imul) Math.imul = function(opA, opB) {
        opB |= 0;
        var result = (opA & 0x003fffff) * opB;
        if (opA & 0xffc00000) result += (opA & 0xffc00000) * opB |0;
        return result |0;
      };
    /****************************************************************************
    var _cuunid = 'wipo-visitor-uunid=';
    uunid(0);

    function uunid(force){
        if (force || document.cookie.indexOf(_cuunid)===-1){
            var value = navigator.userAgent + Date.now() + Math.random().toString().substring(2,11);
            var cookie = _cuunid + cyrb53(value) + ';expires=Jan 2 2034 00:00:00; path=/; SameSite=Lax; domain=.wipo.int';
            document.cookie = cookie;
        }
    }
    ***************************************************************************/
	function wipoVisitor(agent) {
        var value = agent + Date.now() + Math.random().toString().substring(2,11);
        return cyrb53(value, 0);
    }
    function cyrb53(str, seed) {
        seed = seed || 0;
        let h1 = 0xdeadbeef ^ seed, h2 = 0x8badf00d ^ seed;
        for (let i = 0, ch; i < str.length; i++) {
            ch = str.charCodeAt(i);
            h1 = Math.imul(h1 ^ ch, 2654435761);
            h2 = Math.imul(h2 ^ ch, 1597334677);
        }
        h1 = Math.imul(h1 ^ h1>>>16, 2246822507) ^ Math.imul(h2 ^ h2>>>13, 3266489909);
        h2 = Math.imul(h2 ^ h2>>>16, 2246822507) ^ Math.imul(h1 ^ h1>>>13, 3266489909);
        // return 4294967296 * (2097151 & h2) + (h1>>>0);
        return (h2>>>0).toString(16)+(h1>>>0).toString(16);
    }`

//获取js生成的cookie信息
func wipoVisitorUunId(agent string) (string, error) {
	vm     := goja.New()
	_, err := vm.RunString(js)
	if err != nil {
		return "", err
	}
	var wipoVisitor func(string) string
	err     = vm.ExportTo(vm.Get("wipoVisitor"), &wipoVisitor)
	if err != nil {
		return "", err
	}
	return wipoVisitor(agent), nil
}