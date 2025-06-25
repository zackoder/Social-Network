FROM node:20.19.2

WORKDIR /app

RUN npm install -g npm@11.4.2

# Create a non-root user
RUN groupadd -r nodeuser && useradd -r -g nodeuser nodeuser

# Copy package files and install dependencies
COPY package*.json ./
RUN npm install

# Now copy the rest of the app (src/, public/, etc.)
COPY . .
# 11.4.2

# Change ownership to nodeuser
RUN chown -R nodeuser:nodeuser /app

# Switch to non-root user
USER nodeuser

EXPOSE 3000
# RUN next build
# Start the Next.js development server
CMD ["npm", "run", "dev"]