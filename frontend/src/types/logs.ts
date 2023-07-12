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

export type APIUserLog = UserLog;
export type APILogsResponse = Response;