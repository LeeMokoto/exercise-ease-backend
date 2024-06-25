#s3 bucket for RAG data
resource "aws_s3_bucket" "rag_bucket" {
  bucket = "mojima-terraform-state"
    force_destroy = true
  tags = {
    Name        = "Terraform-Bucket"
    Environment = "Prod"
  }
}









