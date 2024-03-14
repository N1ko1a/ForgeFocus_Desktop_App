# ForgeFocus

ForgeFocus is a desktop Electron application designed for Linux that integrates React, Express, and MongoDB to offer a suite of productivity tools including a calendar, to-do list, and Pomodoro timer. This versatile application is tailored to streamline time and task management for enhanced productivity.

[video.webm](https://github.com/N1ko1a/Productivity_Desktop_App/assets/85966654/17d357f4-aa84-440c-b084-5b3614f6264a)
## Features

### Pomodoro Timer
The Pomodoro timer feature helps users boost productivity by automating work and rest intervals. Upon initiation, the timer automatically switches between work and rest sessions, with a default of 4 work sessions followed by breaks. Users receive notifications at the end of each work and rest session, along with updates on the current session progress.

### Calendar
ForgeFocus includes a dynamic calendar feature that notifies users 30 minutes prior to scheduled events. Notifications are also sent at the beginning of each event, ensuring users stay on track with their schedules.

### To-Do List
The to-do list feature is enhanced with workspaces, allowing users to organize tasks efficiently. Users can create, edit, and delete tasks within specific workspaces, ensuring seamless task management.

## Requirements

- Node.js
- MongoDB

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/your-account/ForgeFocus.git
   ```

2. Install the necessary dependencies using npm:

   ```bash
   cd ForgeFocus
   npm install
   ```

3. Build the application for Linux:

   ```bash
   npm run build:linux
   ```

## Running

To run the application execute the following command in dist directory:

```bash
./ForgeFocus-1.0.0.AppImage
```

This will start the Express server for the backend and the Electron application for the frontend.

