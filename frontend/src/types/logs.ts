interface Response {
    error : string;
    messages : UserLog[];
    periods : string[];
}

interface UserLog {
    date : string;
    message : string;
    badges : string;
}

interface ChannelInfo {
    username : string;
    stvID : string;
}

export type APIUserLog = UserLog;
export type APILogsResponse = Response;
export type APIChannelInfo = ChannelInfo;
