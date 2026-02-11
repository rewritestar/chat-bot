"use client";

import { useState } from "react";
import { Chat } from "./_components/chat";
import { MessageType } from "./_components/type";

export default function Room() {
  const [messages, setMessages] = useState<MessageType[]>([]);
  const [input, setInput] = useState("");

  const sendMessage = () => {
    if (input.trim()) {
      const req: MessageType = {
        senderId: "userId",
        content: input,
      };
      setMessages((prev) => [...prev, req]);
      setInput("");
    }
  };

  return (
    <div className="bg-white w-3/5 h-4/5 rounded-xl shadow-xl overflow-hidden flex flex-col">
      <div className="flex">
        <div className="flex items-center">
          <button className="text-2xl text-gray-400 px-4 border-r-2 border-r-gray-100">{`<`}</button>
        </div>
        <div className="border-b-2 border-b-gray-100 p-4">
          <span>방이름</span>
        </div>
      </div>
      <div className="flex-1 flex flex-col min-h-0 inset-shadow-sm">
        <div className="flex-1 overflow-auto p-4">
          {messages.map((message, i) => (
            <Chat key={i} data={message} myId={"userId"} />
          ))}
        </div>
        <div className="h-16 bg-blue-50 flex items-center p-2">
          <textarea
            className="bg-white rounded-md shadow-md flex-1 mr-2 h-10 resize-none focus:outline-2 focus:outline-blue-200 p-2"
            value={input}
            onChange={(e) => setInput(e.target.value)}
          />
          <input
            type="button"
            value="Send"
            onClick={sendMessage}
            className="bg-blue-400 rounded-md text-white w-16 h-10"
          />
        </div>
      </div>
    </div>
  );
}
