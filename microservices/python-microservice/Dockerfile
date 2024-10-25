# Start with the base image
FROM python:3.12.7-alpine3.20 

# # Set a working directory
# WORKDIR /app-dir

# Copy the requirements and install dependencies
COPY ./app/requirements.txt .
RUN pip3 install --no-cache-dir --upgrade -r requirements.txt 

# Copy the rest of the app's code
COPY . .

# Create a non-root user and group
RUN adduser -D nonrootuser

# Give non-root user ownership of the app directory
RUN chown -R nonrootuser:nonrootuser /app

# Switch to the non-root user
USER nonrootuser

# Expose the port
EXPOSE 7777

RUN cd app

# Run the FastAPI app with Uvicorn as a non-root user
CMD ["uvicorn", "app:app", "--host", "0.0.0.0", "--port", "7777"]


