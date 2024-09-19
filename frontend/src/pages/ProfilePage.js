import React, { useState, useEffect } from 'react';
import { getUserProfile } from '../services/authservice';

const ProfilePage = () => {
  const [user, setUser] = useState(null);

  useEffect(() => {
    const fetchUserProfile = async () => {
      const data = await getUserProfile();
      setUser(data);
    };

    fetchUserProfile();
  }, []);

  if (!user) return <p>Loading...</p>;

  return (
    <div>
      <h1>Profile</h1>
      <p>Username: {user.username}</p>
    </div>
  );
};

export default ProfilePage;

