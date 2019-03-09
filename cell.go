package main

type cell struct {
	uptime           int64
	compressed       bool
	refractor        bool
	compressedTime   int64
	refractorTime    int64
	compressedSize   int64
	decompressedSize int64
	dead             bool
	delta            int64
	potential        bool
}

func (c *cell) GivePotential (timer int64) {
	if !c.dead && !c.compressed {
		c.potential = true
		c.compressed = true
		c.uptime = timer
	}
}

func (c *cell) Decompress () {
	c.refractor = false
	c.compressed = false
}

func (c *cell) SetRefractor () {
	c.refractor = true
	c.potential = false
}

func Kill () *cell {
	return &cell{
		0,
		true,
		true,
		0,
		0,
		0,
		0,
		true,
		0,
		false,
	}
}