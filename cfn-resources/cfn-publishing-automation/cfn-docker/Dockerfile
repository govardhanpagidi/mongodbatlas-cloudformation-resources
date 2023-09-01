# Use an official Amazon Linux 2 image as the base
FROM amazonlinux:2

# Install essential tools
RUN yum update -y && \
yum install -y python3 && \
yum install -y golang && \
yum install -y unzip && \
curl "https://d1vvhvl2y92vvt.cloudfront.net/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip" && \
unzip awscliv2.zip && \
./aws/install && \
rm -rf awscliv2.zip aws

# Install AWS SAM CLI
RUN curl -L "https://github.com/aws/aws-sam-cli/releases/latest/download/aws-sam-cli-linux-x86_64.zip" -o /tmp/sam.zip && \
unzip -d /usr/local/bin /tmp/sam.zip && \
rm /tmp/sam.zip

# Set environment variables for Golang
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH


# Copy the script into the image
COPY cfn-install-mongo-cli.sh /cfn-install-mongo-cli.sh

RUN echo "Installing MongoDB CLI"
# Make the script executable and run it
RUN chmod +x /cfn-install-mongo-cli.sh

RUN /cfn-install-mongo-cli.sh

