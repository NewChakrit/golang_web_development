import React, { useEffect, useRef, useState } from "react";
import { useNavigate } from "react-router-dom";
import Header from "./Header";
import "./Dashboard.css";

const Dashboard = () => {
  const [userInfo, setUserInfo] = useState(null);
  const [isLoading, setisLoading] = useState(true);
  const navigate = useNavigate();
  const hasFetched = useRef(false); // ✅ Tracks API call execution

  useEffect(() => {
    const fetchUserInfo = async () => {
      try {
        if (hasFetched.current) return; // ✅ Prevents duplicate API call
        hasFetched.current = true; // ✅ Mark API as called

        const queryParams = new URLSearchParams(window.location.search);
        const state = queryParams.get("state");
        const code = queryParams.get("code");

        if (state && code) {
          setisLoading(true)
          const response = await fetch(
            `http://localhost:8080/callback/google?state=${state}&code=${code}`
          );
          if(!response.ok){
            throw new Error("error")
          }
          const data = await response.json();
          setUserInfo(data);
          console.log("Picture URL:", userInfo?.picture);
        }
      } catch (error) {
        console.error("Error fetching user info:", error);
        navigate("/")
      }finally{
        setisLoading(false)
      }
    };

    fetchUserInfo();
  }, []);

  if (isLoading) {
    return (
      <div className="dashboard-container">
        <div className="loading-spinner">Loading...</div>
      </div>
    );
  }

  return (
    <>
      <Header />
      <div className="dashboard-container">
        <div className="dashboard-content">
          <div className="profile-section">
            <img 
              src={userInfo?.picture || "https://via.placeholder.com/100"}
              alt="Profile"
              className="profile-image"
            />
            <h1>Welcome, {userInfo?.name}</h1>
            <p className="email">{userInfo?.email}</p>
          </div>
          <button className="logout-button" onClick={() => navigate("/")}>
            Logout
          </button>
        </div>
      </div>
    </>
  );
};

export default Dashboard;
