# Use an official node image as the base
FROM node:18-alpine AS build

# Set working directory
WORKDIR /app

# Copy package.json and package-lock.json
COPY package.json package-lock.json ./

# Install dependencies
RUN npm install

# Copy the rest of the app files
COPY . .

# Build the app
RUN npm run build

# Serve the build with a lightweight server (e.g., serve)
# Install a simple HTTP server to serve the static files
RUN npm install -g serve

# Expose port 5000 to the outside world
EXPOSE 5000

# Start the app by serving the static files
CMD ["serve", "-s", "dist", "-l", "5000"]

