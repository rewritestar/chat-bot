"use client";

import { useRouter } from "next/navigation";

export default function Home() {
  const naviRouter = useRouter();
  naviRouter.push("/room");
  return <></>;
}
