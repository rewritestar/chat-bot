import { ChatProps } from "./type";

export function Chat({ data, myId }: ChatProps) {
  const isMyChat = data.senderId === myId;
  const chatBg = isMyChat ? "bg-blue-50" : "bg-blue-300";
  const chatPosition = isMyChat ? "justify-end" : "justify-start";
  return (
    <div className={`w-full flex ${chatPosition} mb-2`}>
      <div className={`${chatBg} rounded-md shadow-md p-2 w-fit max-w-2/5`}>
        {data.content}
      </div>
    </div>
  );
}
