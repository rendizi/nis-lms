import React, { useState, useEffect, useRef } from 'react';
import Navbar from '../Components/Navbar';

function Top() {
  const [users, setUsers] = useState([]);
  const [selectedUser, setSelectedUser] = useState(null);
  const dialogRef = useRef(null);

  useEffect(() => {
    fetch('http://localhost:8080/top')
      .then(response => response.json())
      .then(data => setUsers(data))
      .catch(error => console.error('There was an error fetching the data!', error));
  }, []);

  const handleUserClick = (user) => {
    setSelectedUser(user);
    if (dialogRef.current) {
      dialogRef.current.showModal();
    }
  };

  const closeModal = () => {
    if (dialogRef.current) {
      dialogRef.current.close();
    }
    setSelectedUser(null);
  };

  const renderTrophy = (index) => {
    switch (index) {
      case 0:
        return '🏆';
      case 1:
        return '🥈';
      case 2:
        return '🥉';
      default:
        return '';
    }
  };

  return (
    <>
    <Navbar />
    <div className="flex flex-col items-center bg-base-100 p-4">
      <h1 className="text-4xl font-bold mb-8">Top Users</h1>
      <div className="flex flex-wrap justify-center gap-8">
        {users.map((user, index) => (
          <div
            key={user.id}
            onClick={() => handleUserClick(user)}
            className="bg-base-content p-6 rounded-lg shadow-lg cursor-pointer hover:scale-105 transform transition"
          >
            {index < 3 && <span className="text-2xl">{renderTrophy(index)}</span>}
            <p className="text-xl font-semibold mt-2 text-white"><strong>Username:</strong> {user.login}</p>
            <p className="mt-2 text-white"><strong>Solved:</strong> {user.stats.solved}</p>
            <p className="mt-2 text-white"><strong>Rating:</strong> {user.stats.rating}</p>
          </div>
        ))}
      </div>
      <dialog ref={dialogRef} className="bg-base-content rounded-lg p-6 shadow-lg max-w-md w-full">
        {selectedUser && (
          <>
            <h2 className="text-2xl font-bold mb-4">{selectedUser.id}. {selectedUser.login}</h2>
            <p><strong>Email:</strong> {selectedUser.email}</p>
            <p><strong>Class:</strong> {selectedUser.klass}</p>
            <p><strong>School:</strong> {selectedUser.school}</p>
            <p><strong>Solved:</strong> {selectedUser.stats.solved}</p>
            <p><strong>Leetcode:</strong> {selectedUser.stats.leetcode}</p>
            <p><strong>Badges:</strong> {selectedUser.stats.badges}</p>
            <p><strong>Rating:</strong> {selectedUser.stats.rating}</p>
            <button onClick={closeModal} className="mt-4 px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-700">
              Close
            </button>
          </>
        )}
      </dialog>
    </div>
    </>
  );
}

export default Top;
