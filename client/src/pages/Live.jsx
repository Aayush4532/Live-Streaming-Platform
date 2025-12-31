import React, { useEffect, useRef, useState } from "react";

const Live = () => {
  const videoRef = useRef(null);
  const refreshTimerRef = useRef(null);

  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");

  const PLAYBACK_API = "http://localhost:1102/api/user/join-seminar";

  const fetchPlayback = async (isInitial = false) => {
    try {
      const res = await fetch(PLAYBACK_API);
      if (!res.ok) throw new Error("Stream unavailable");

      const data = await res.json();
      const playbackUrl = data.playback.url;
      const refreshIn = data.refresh_in;

      if (videoRef.current) {
        videoRef.current.src = playbackUrl;
      }

      if (refreshTimerRef.current) {
        clearTimeout(refreshTimerRef.current);
      }

      refreshTimerRef.current = setTimeout(() => {
        fetchPlayback(false);
      }, Math.max((refreshIn - 10) * 1000, 10000));

      setLoading(false);
    } catch (err) {
      setError("Live stream is not available right now");
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchPlayback(true);

    return () => {
      if (refreshTimerRef.current) {
        clearTimeout(refreshTimerRef.current);
      }
    };
  }, []);


  if (loading) {
    return (
      <div className="h-screen w-screen bg-black flex items-center justify-center">
        <p className="text-zinc-400 text-sm">Connecting to live stream…</p>
      </div>
    );
  }

  if (error) {
    return (
      <div className="h-screen w-screen bg-black flex items-center justify-center">
        <div className="text-center">
          <p className="text-white text-lg">Stream Offline</p>
          <p className="text-zinc-400 text-sm">{error}</p>
        </div>
      </div>
    );
  }

  return (
    <div className="h-screen w-screen bg-black flex items-center justify-center">
      <div className="w-full max-w-screen h-auto aspect-video flex items-center justify-center">
        <video
          ref={videoRef}
          controls
          autoPlay
          playsInline
          className="w-full h-full border-b bg-black object-contain"
        />
      </div>

    </div>
  );
};

export default Live;
