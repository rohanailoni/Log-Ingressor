CREATE TABLE ErrorLog (
                     id SERIAL PRIMARY KEY,
                     level VARCHAR(255),
                     message TEXT,
                     resourceId VARCHAR(255),
                     timestamp TIMESTAMP,
                     traceId VARCHAR(255),
                     spanId VARCHAR(255),
                     commit VARCHAR(255),
                     parentResourceId VARCHAR(255)
);

-- Create an index on timestamp and resourceId
CREATE INDEX idx_log_timestamp_resourceId ON ErrorLog (timestamp, resourceId);

-- Insert data into the Log table
CREATE TABLE DebugLog (
                    id SERIAL PRIMARY KEY,
                     level VARCHAR(255),
                     message TEXT,
                     resourceId VARCHAR(255),
                     timestamp TIMESTAMP,
                     traceId VARCHAR(255),
                     spanId VARCHAR(255),
                     commit VARCHAR(255),
                     parentResourceId VARCHAR(255)
);

-- Create an index on timestamp and resourceId
CREATE INDEX idx_log_timestamp_resourceId ON DebugLog (timestamp, resourceId);