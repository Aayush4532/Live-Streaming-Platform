import React from "react";
import { useNavigate } from "react-router-dom";

const Home = () => {
  const navigate = useNavigate();

  return (
    <div className="h-screen w-screen flex items-center justify-center bg-black">
      <div className="text-center">

        <h1 className="text-white text-2xl font-medium mb-2">
          Live Seminar Platform
        </h1>

        <p className="text-zinc-400 text-sm mb-8">
          Join live sessions or host your own seminar
        </p>

        <div className="flex gap-4 justify-center">
          <button
            onClick={() => navigate("/live-seminar")}
            className="px-6 py-2 border border-zinc-700 text-white rounded-md"
          >
            Join Seminar
          </button>

          <button
            onClick={() => navigate("/login")}
            className="px-6 py-2 border border-zinc-700 text-white rounded-md"
          >
            Host Login
          </button>
        </div>

      </div>
    </div>
  );
};

export default Home;
