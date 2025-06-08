// @refresh reload
import { mount, StartClient } from "@solidjs/start/client";

//function sleep(ms: number) {
//    return new Promise(res => setTimeout(res, ms))
//}
//
//await sleep(2000)
//console.log("Entry")

mount(() => <StartClient />, document.getElementById("app")!);
