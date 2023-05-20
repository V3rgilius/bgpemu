
class BgpUpdateElem():
    def __init__(self,raw,prefix,peer_asn,peer_ip) -> None:
        self.raw = raw
        self.prefix = prefix
        self.peer_as = peer_asn
        self.peer_ip = peer_ip