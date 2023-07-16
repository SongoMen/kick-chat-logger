
interface UserInfo {
    connections: Connection[];
    name: string;
    emotes: Emote[];
}

interface Connection {
    id: string;
    platform: string;
    emote_set: EmoteSet;
}

interface EmoteSet {
    id: string;
    name: string;
    emotes: Emote[];
}

interface Emote {
    id: string;
    name: string;
    data: EmoteData;
}

interface EmoteData {
    id: string;
    name: string;
    listed: boolean;
    animated: boolean;
    host: EmoteHost;
}

interface EmoteHost {
    url: string;
    files: EmoteFile[];
}

interface EmoteFile {
    name: string;
    width: number;
    height: number;
    format: string;
}

export type SvtEmoteSet = EmoteSet
export type SvtEmote = Emote
export type SvtUserInfo = UserInfo
export type SvtConnection = Connection
