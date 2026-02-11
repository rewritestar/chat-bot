export type MessageType = {
  senderId: string;
  content: string;
};

export type ChatProps = {
  data: MessageType;
  myId: string;
};
