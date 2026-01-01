import React, { useEffect, useRef, useState } from "react";
import Hls from "hls.js";

const Live = () => {
  const videoRef = useRef(null);
  const hlsRef = useRef(null); 
  const refreshTimerRef = useRef(null);

  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");

  const PLAYBACK_API =
    "https://live-streaming-platform-production.up.railway.app/api/user/join-seminar";

  const attachStream = (playbackUrl) => {
    const video = videoRef.current;
    if (!video) return;

    if (hlsRef.current) {
      hlsRef.current.destroy();
      hlsRef.current = null;
    }

    if (video.canPlayType("application/vnd.apple.mpegurl")) {
      video.src = playbackUrl;
      video.play();
      return;
    }

    if (Hls.isSupported()) {
      const hls = new Hls();
      hls.loadSource(playbackUrl);
      hls.attachMedia(video);
      hlsRef.current = hls;
    } else {
      setError("Browser does not support HLS");
    }
  };

  const fetchPlayback = async () => {
    try {
      const res = await fetch(PLAYBACK_API);
      if (!res.ok) throw new Error("Stream unavailable");

      const data = await res.json();
      attachStream(data.playback.url);

      setLoading(false);
    } catch (err) {
      setError("Live stream is not available right now");
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchPlayback();

    return () => {
      if (refreshTimerRef.current) clearTimeout(refreshTimerRef.current);
      if (hlsRef.current) hlsRef.current.destroy();
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
        <p className="text-white">{error}</p>
      </div>
    );
  }

  return (
    <div className="h-screen w-screen bg-black flex items-center justify-center">
      <video
        ref={videoRef}
        controls
        autoPlay
        playsInline
        className="w-full h-full bg-black object-contain"
      />
    </div>
  );
};

export default Live;
