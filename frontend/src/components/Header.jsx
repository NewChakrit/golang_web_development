import React from 'react';
import './Header.css';

const Header = () => {
  const handleGithubClick = () => {
    window.open('https://github.com/rajanlagah/react-login', '_blank');
  };

  return (
    <header className="header">
      <div className="header-content">
        <div className="github-link" onClick={handleGithubClick}>
          <img 
            src="https://github.githubassets.com/images/modules/logos_page/GitHub-Mark.png"
            alt="GitHub Logo" 
            className="github-logo"
          />
          <span className="github-text">View on GitHub</span>
        </div>
      </div>
    </header>
  );
};

export default Header; 