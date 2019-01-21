# config
support read some type of config files,current support yaml and json

example:


type Configger  struct {
	zk string
	kafka string
	srcGroup string
	log_path string
}
func (c *Configger) GetZk()(zk string ){
	zk=c.zk
	return zk
}

func (c *Configger) GetKafka()(kafka string ){
	kafka=c.kafka
	return kafka
}

func (c *Configger) GetSrcGroup()(srcGroup string ){
	srcGroup=c.srcGroup
	return srcGroup
}

func (c *Configger) GetLog_path()(log_path string ){
	log_path=c.log_path
	return log_path
}

func (c *Configger ) Init(filenmae string) (error){

	err:=config.Parse(filenmae,c)
	return err

}
func main(){
  c:=flag.String("conf","conf.yml","config info from cofigFile")
	flag.Parse()

	confS:=new(Configger)

	confS.Init(*c)


}

