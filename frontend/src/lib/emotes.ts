import { getChannelSvtEmotes, getGlobalSvtEmotes } from "./7tv"

const getAllEmotes = async (stvID: string): Promise<Record<string, string>> => {
    const emoteProviders = [getChannelSvtEmotes, getGlobalSvtEmotes];
    const allEmotes: (void | Record<string, string>)[] = await Promise.all(emoteProviders.map(provider => provider(stvID)));
    const mergedEmotes: Record<string, string> = {};
    allEmotes.forEach(emotes => {
        if (emotes) {
            Object.entries(emotes).forEach(([key, value]) => {
                mergedEmotes[key] = value;
            });
        }
    });
    return mergedEmotes;
}

export {
    getAllEmotes
}