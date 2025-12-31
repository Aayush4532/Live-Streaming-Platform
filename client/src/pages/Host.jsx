import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

const Host = () => {
    const [loading, setLoading] = useState(true);
    const [url, setUrl] = useState("");
    const [key, setKey] = useState("");
    const [showKey, setShowKey] = useState(false);
    const navigate = useNavigate();

    useEffect(() => {
        const checkAuth = async () => {
            try {
                const res = await fetch("https://live-streaming-platform-production.up.railway.app/api/host/check", {
                    method: "GET",
                    credentials: "include",
                });

                if (!res.ok) {
                    navigate("/login");
                    return;
                }

                setLoading(false);
            } catch (err) {
                navigate("/login");
            }
        };

        checkAuth();
    }, [navigate]);

    const handleStartSeminar = async () => {
        try {
            const res = await fetch("https://live-streaming-platform-production.up.railway.app/api/host/create", {
                method: "GET",
                credentials: "include",
            });

            if (!res.ok) {
                alert("Session couldn't be created");
                return;
            }

            const data = await res.json();
            setUrl(data.rtmp_uri);
            setKey(data.stream_key);
        } catch (err) {
            alert("Something went wrong");
        }
    };

    const copyToClipboard = (text) => {
        navigator.clipboard.writeText(text);
        alert("Copied to clipboard");
    };

    if (loading) {
        return (
            <div className="h-screen w-screen flex items-center justify-center bg-black">
                <p className="text-zinc-400 text-sm">Checking authentication...</p>
            </div>
        );
    }

    return (
        <div className="h-screen w-screen flex items-center justify-center bg-black">
            <div className="w-full max-w-md border border-zinc-800 rounded-lg p-6 bg-zinc-950">

                <h1 className="text-white text-xl font-medium mb-6 text-center">
                    Host Dashboard
                </h1>

                {!url && !key && (
                    <button
                        className="w-full py-2 border border-zinc-700 text-white rounded-md"
                        onClick={handleStartSeminar}
                    >
                        Start Seminar
                    </button>
                )}

                {url && key && (
                    <div className="space-y-4">

                        {/* RTMP URL */}
                        <div>
                            <p className="text-zinc-400 text-sm mb-1">RTMP URL</p>
                            <div className="flex gap-2">
                                <input
                                    value={url}
                                    readOnly
                                    className="flex-1 px-3 py-2 bg-black text-white border border-zinc-700 rounded-md text-sm"
                                />
                                <button
                                    className="px-3 border border-zinc-700 text-white rounded-md"
                                    onClick={() => copyToClipboard(url)}
                                >
                                    Copy
                                </button>
                            </div>
                        </div>

                        {/* Stream Key */}
                        <div>
                            <p className="text-zinc-400 text-sm mb-1">Stream Key</p>
                            <div className="flex gap-2">
                                <input
                                    value={showKey ? key : "••••••••••••••••"}
                                    readOnly
                                    className="flex-1 px-3 py-2 bg-black text-white border border-zinc-700 rounded-md text-sm"
                                />
                                <button
                                    className="px-3 border border-zinc-700 text-white rounded-md"
                                    onClick={() => setShowKey(!showKey)}
                                >
                                    {showKey ? "Hide" : "Show"}
                                </button>
                                <button
                                    className="px-3 border border-zinc-700 text-white rounded-md"
                                    onClick={() => copyToClipboard(key)}
                                >
                                    Copy
                                </button>
                            </div>
                        </div>

                    </div>
                )}

            </div>
        </div>
    );
};

export default Host;
