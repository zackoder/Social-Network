FROM node:20.19.2

WORKDIR /app

# Create a non-root user
RUN groupadd -r nodeuser && useradd -r -g nodeuser nodeuser

# Copy package files
COPY package*.json ./

# Install dependencies
RUN npm install

# Change ownership to nodeuser
RUN chown -R nodeuser:nodeuser /app

# Switch to non-root user
USER nodeuser

# Expose port
EXPOSE 3000

# Default command for development
CMD ["npm", "run", "dev"]