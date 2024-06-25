#s3 bucket for RAG data
resource "aws_s3_bucket" "ee_bucket" {
  bucket = "ee-artefacts-bucket"
    force_destroy = true
  tags = {
    Name        = "EE-Bucket"
    Environment = "Prod"
  }
}

resource "aws_s3_bucket_cors_configuration" "example" {
  bucket = aws_s3_bucket.ee_bucket.id

  cors_rule {
    
    allowed_headers = ["*"]
    allowed_methods = ["PUT", "POST", "HEAD"]
    allowed_origins = ["*"]
    expose_headers  = ["ETag"]
    max_age_seconds = 6000
  }

}





