import boto3

# AWS 자격 증명을 로드하고 세션을 생성합니다.
session = boto3.Session(
    aws_access_key_id='',
    aws_secret_access_key='',
    region_name='ap-northeast-2'
)

# 세션을 사용하여 S3 클라이언트를 초기화합니다.
s3 = session.client('s3')
bucket_name = 'test-guiwoo'
object_key = 'example2.png'

def getImage():
    try:
        # S3 객체 가져오기
        response = s3.get_object(Bucket=bucket_name, Key=object_key)
        # 객체 데이터 읽기
        object_data = response['Body'].read()
        # 가져온 데이터 출력 또는 원하는 작업 수행
        print(object_data)
    except Exception as e:
         print("Error:", e)


def uploadImage():
    local_image_path = '/Users/guiwoopark/Desktop/personal/study/comfyui/assets/example2.png'

    try:
        with open(local_image_path,'rb') as f:
            s3.upload_fileobj(f,bucket_name,"guiwoo_test_upload.png")

        print("Image uploaded successfully.")

    except Exception as e:
        print("Error : ",e)


uploadImage()