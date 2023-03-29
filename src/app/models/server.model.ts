export interface Server {
    id: string;
    name: string;
    admin: string;
    members: string[];
    messages: ServerMessage[];
}
  
export interface ServerMessage {
    username: string;
    message: string;
    timestamp: Date;
}
  