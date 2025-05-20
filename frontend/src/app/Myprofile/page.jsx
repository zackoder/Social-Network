"use client";

import { useEffect } from "react";
import { useRouter, useSearchParams } from "next/navigation";
import ProfilePage from "../profile/page";

export default function MyProfilePage() {
  const router = useRouter();
  const searchParams = useSearchParams();
  const userId = searchParams.get("id");

  useEffect(() => {
    if (!userId) {
      // If no ID is provided, redirect back to home
      router.push("/");
    }
  }, [userId, router]);

  if (!userId) {
    return <div>Loading profile...</div>;
  }

  // Pass the userId to the existing ProfilePage component
  return <ProfilePage params={{ id: userId }} />;
}
