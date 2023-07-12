import type { SvtEmote } from "$types/7tv";

const getGlobalEmotes = async () => {
    const req = await fetch("https://7tv.io/v3/emote-sets/global");
    const data = await req.json();
    return retrieveEmotes(data.emotes);
};

const getChannelEmotes = async (channel: string) => {

};

const retrieveEmotes = async (emotesArray: SvtEmote[]) => {
    const emotes = [];
    for (let i = 0; i < emotesArray.length; i++) {
        const webpEmotes = emotesArray[i].data.host.files.filter(i => i.format === 'WEBP');
        const emoteURL = emotesArray[i].data.host.url;
        emotes.push({
            name: emotesArray[i].name,
            url: emoteURL + "/" + webpEmotes[0].name,
        });
    }
    return emotes;
};

export {
    getGlobalEmotes,
    getChannelEmotes
}