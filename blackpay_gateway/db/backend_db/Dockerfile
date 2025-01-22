FROM cassandra:latest

# Copy custom configuration files if needed
# COPY ./conf/cassandra.yaml /etc/cassandra/cassandra.yaml

# Expose necessary ports
EXPOSE 7000 7001 7199 9042 9160

# Run the Cassandra server
CMD ["cassandra", "-f"]
