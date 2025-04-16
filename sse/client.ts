import { parseEventData, uint8ToString} from './util.ts';

interface Data {
    name: string;
    count: number;
}

const resp = await fetch("http://localhost:8888/sse", { method: "POST"});
if (resp.body === null) {
    console.log(resp.status)
    Deno.exit(0)
}

for await (const x of resp.body) {
    const data = parseEventData<Data>(await uint8ToString(x))
    if (data === null) {
        continue
    }
    if (data.count > 20) {
        break;
    }
    console.log(data.name, data.count)
}
resp.body.cancel()
