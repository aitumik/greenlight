# Use official Bun image
FROM oven/bun:latest

# Set working directory inside the container
WORKDIR /app

# Copy package files
COPY bun.lock package.json vite.config.js ./

# Install dependencies using Bun
RUN bun install

# Copy the rest of your project
COPY . .

# Expose Vite's default port
EXPOSE 5173

# Start the Vite dev server
CMD ["bun", "run", "dev"]

