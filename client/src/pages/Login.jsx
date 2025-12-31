import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

const Login = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [loading, setLoading] = useState(false);
  const [checkingAuth, setCheckingAuth] = useState(true);

  const navigate = useNavigate();
  useEffect(() => {
    const checkAuth = async () => {
      try {
        const res = await fetch(
          "https://live-streaming-platform-production.up.railway.app/api/host/check",
          {
            method: "GET",
            credentials: "include",
          }
        );

        if (res.ok) {
          navigate("/host");
          return;
        }
      } catch (err) {
        console.error("auth check failed:", err);
      } finally {
        setCheckingAuth(false);
      }
    };

    checkAuth();
  }, [navigate]);

  const handleSubmitClick = async () => {
    if (!email || !password) {
      alert("Email and password are required");
      return;
    }

    setLoading(true);

    try {
      const res = await fetch(
        "https://live-streaming-platform-production.up.railway.app/api/auth/login",
        {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          credentials: "include",
          body: JSON.stringify({ email, password }),
        }
      );

      if (!res.ok) {
        const msg = await res.text();
        alert(msg || "Invalid credentials");
        return;
      }

      // cookie set ho chuki hai
      navigate("/host");
    } catch (err) {
      alert("Something went wrong. Try again.");
    } finally {
      setLoading(false);
    }
  };

  // ⏳ While checking existing session
  if (checkingAuth) {
    return (
      <div className="h-screen w-screen flex items-center justify-center bg-black">
        <p className="text-zinc-400 text-sm">Checking session...</p>
      </div>
    );
  }

  return (
    <div className="h-screen w-screen flex items-center justify-center bg-black">
      <div className="w-full max-w-sm border border-zinc-800 rounded-lg p-6 bg-zinc-950">

        <h1 className="text-white text-xl font-medium mb-6 text-center">
          Host Login
        </h1>

        <input
          type="email"
          placeholder="Email"
          className="w-full mb-4 px-3 py-2 bg-black text-white border border-zinc-700 rounded-md focus:outline-none"
          onChange={(e) => setEmail(e.target.value)}
        />

        <input
          type="password"
          placeholder="Password"
          className="w-full mb-6 px-3 py-2 bg-black text-white border border-zinc-700 rounded-md focus:outline-none"
          onChange={(e) => setPassword(e.target.value)}
        />

        <button
          onClick={handleSubmitClick}
          disabled={loading}
          className="w-full py-2 border border-zinc-700 text-white rounded-md"
        >
          {loading ? "Please wait..." : "Login"}
        </button>

      </div>
    </div>
  );
};

export default Login;
