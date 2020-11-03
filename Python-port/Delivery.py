import requests as req
class Delivery:
    def __init__(self, name="", address="", time="", content="", special=False, comment=""):
        self.name = name
        # String, lat, lng
        self.address = address
        self.lat_lng = (None, None)
        self.time = time
        self.content = content
        self.special = special
        self.comment = comment
        self.failed = True
        self.group = None


    

    def set_group(self, grp):
        self.group = grp


    
    def __str__(self):
        # Returns geojson format
        return ""
