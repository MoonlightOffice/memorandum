export function parseEventData<T>(data: string): T | null {
    if (!data.startsWith('data:')) {
        return null;
    }

    return JSON.parse(data.trim().substring(6))
}

export function sleep(ms: number) {
    return new Promise(res => setTimeout(res, ms))
}

export async function uint8ToString(b: Uint8Array): Promise<string> {
    return await new Blob([b]).text()
}
