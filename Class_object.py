# Your solution here
class youtubechannel():
  total = 0
  def __init__(self, channelname, channeltype):
    self.channelname = channelname
    self.channeltype = channeltype
  
  def subscribe(self):
    youtubechannel.total = youtubechannel.total + 1
    print(f"Thanks for susbscribing. Total subcribers:{youtubechannel.total},{self.channelname} and {self.channeltype}")

  def unsubscribe(self):
    youtubechannel.total = youtubechannel.total - 1
    print(f"Thanks for unsusbscribing. Total subcribers,{youtubechannel.total}")

  def show(self):
    print(f"Existing susbscribers. Total subcribers, {youtubechannel.total}")

  @classmethod
  def reset_total(cls):
    cls.total = 0







codebasics = youtubechannel("codebasics", "news")
codebasics.subscribe()
codebasics.show()

xyz = youtubechannel("funny", "24X7")
xyz.subscribe()
xyz.show()

#xyz.unsubscribe()

abc = youtubechannel("StudyContent", "24X7")
abc.subscribe()
abc.show()

youtubechannel.reset_total()