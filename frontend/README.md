# React Google OAuth Login

A simple and clean React application demonstrating Google OAuth 2.0 authentication implementation. This project is part of my Udemy course on "Modern Authentication with React".

## Overview

This project showcases how to implement Google OAuth 2.0 authentication in a React application, featuring:
- A landing page with Google login button
- Protected dashboard route
- User profile display after successful authentication
- Clean and responsive UI

## Features

- 🔐 Google OAuth 2.0 Authentication
- 👤 User Profile Display
- 🎨 Clean, Modern UI
- 🔒 Protected Routes
- 🔄 Loading States
- 📱 Responsive Design

## Project Structure
```
src/
├── components/
│ ├── Dashboard.jsx # Protected dashboard showing user profile
│ ├── Landing.jsx # Landing page with Google login
│ └── Header.jsx # Common header component
└── App.js # Main routing component
```

## Installation

1. Clone the repository:

```bash
git clone https://github.com/yourusername/react-google-oauth.git
cd react-google-oauth
```

2. Install dependencies:

```bash
npm install
```

3. Start the development server:

```bash
npm start
```

## Usage

1. Open your browser and navigate to `http://localhost:3000`
2. Click the Google login button to authenticate (make sure go server is running on port 8080)
3. Access the protected dashboard route
