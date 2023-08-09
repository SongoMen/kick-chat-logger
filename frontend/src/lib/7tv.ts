import type { SvtEmote, SvtUserInfo, SvtConnection } from "$types/7tv";

const getGlobalEmotes = async (): Promise<Record<string, string>> => {
    const req = await fetch("https://7tv.io/v3/emote-sets/global");
    const data = await req.json();
    return retrieveEmotes(data.emotes);
};

const getChannelEmotes = async (channel: string): Promise<Record<string, string>> => {
    if (!channel) {
        return {};
    }
    const req = await fetch(`https://7tv.io/v3/users/${channel}`);
    const data: SvtUserInfo = await req.json();
    const connections: Record<string, SvtConnection> = {}
    for (const conenction of (data.connections || [])) {
        connections[conenction.platform] = conenction
    }
    const rightConnection = connections['KICK'] || connections['TWITCH'] || null;
    if (!rightConnection) {
        return {};
    }
    const emoteSetReq = await fetch(`https://7tv.io/v3/emote-sets/${rightConnection.emote_set.id}`);
    const emoteSetData = await emoteSetReq.json();
    return retrieveEmotes(emoteSetData.emotes);
};

const retrieveEmotes = (emotesArray: SvtEmote[]): Record<string, string> => {
    const emotes: Record<string, string> = {};
    for (let i = 0; i < (emotesArray || []).length; i++) {
        const webpEmotes = emotesArray[i].data.host.files.filter(i => i.format === 'WEBP');
        const emoteURL = emotesArray[i].data.host.url;
        emotes[emotesArray[i].name] = emoteURL + "/" + webpEmotes[0].name
    }
    return emotes;
};

export const getGlobalSvtEmotes = getGlobalEmotes;
export const getChannelSvtEmotes = getChannelEmotes;