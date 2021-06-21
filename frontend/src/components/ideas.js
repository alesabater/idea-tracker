import React from 'react'

const Ideas = ({ ideas }) => {
  return (
    <div>
      <center><h1>Idea List</h1></center>
      {ideas.map((idea) => (
        <div class="card">
          <div class="card-body">
            <h5 class="card-title">{idea.name}</h5>
            <p class="card-text">{idea.description}</p>
          </div>
        </div>
      ))}
    </div>
  )
};

export default Ideas