package nats

func (nc *Conn) ConnectedServerHost() string {
	if nc == nil {
		return _EMPTY_
	}
	nc.mu.Lock()
	defer nc.mu.Unlock()
	if nc.status != CONNECTED {
		return _EMPTY_
	}
	return nc.info.Host
}

func (nc *Conn) ConnectedServerVersion() string {
	if nc == nil {
		return _EMPTY_
	}
	nc.mu.Lock()
	defer nc.mu.Unlock()
	if nc.status != CONNECTED {
		return _EMPTY_
	}
	return nc.info.Version
}
