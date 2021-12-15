from diagrams import Cluster, Diagram
from diagrams.programming.language import Go
from diagrams.onprem.database import PostgreSQL
from diagrams.generic.blank import Blank
from diagrams.aws.storage import SimpleStorageServiceS3Bucket

with Diagram("Spudify", show=False):
    with Cluster("Service"):
        backend = Go("API")
        database = PostgreSQL("Database")
        
        backend >> database

    object_storage = SimpleStorageServiceS3Bucket("Object Storage")
    client = Blank("Client")

    client >> backend
    client >> object_storage
