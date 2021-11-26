package world

import (
	"github.com/df-mc/dragonfly/server/world/chunk"
	"io"
)

// Provider represents a value that may provide world data to a World value. It usually does the reading and
// writing of the world data so that the World may use it.
type Provider interface {
	io.Closer
	// Settings returns the settings for a World.
	Settings() Settings
	// SaveSettings saves the settings of a World.
	SaveSettings(Settings)

	// LoadChunk attempts to load a chunk from the chunk position passed. If successful, a non-nil chunk is
	// returned and exists is true and err nil. If no chunk was saved at the chunk position passed, the chunk
	// returned is nil, and so is the error. If the chunk did exist, but if the data was invalid, nil is
	// returned for the chunk and true, with a non-nil error.
	// If exists ends up false, the chunk at the position is instead newly generated by the world.
	LoadChunk(position ChunkPos) (c *chunk.Chunk, exists bool, err error)
	// SaveChunk saves a chunk at a specific position in the provider. If writing was not successful, an error
	// is returned.
	SaveChunk(position ChunkPos, c *chunk.Chunk) error
	// LoadEntities loads all entities stored at a particular chunk position. If the entities cannot be read,
	// LoadEntities returns a non-nil error.
	LoadEntities(position ChunkPos) ([]SaveableEntity, error)
	// SaveEntities saves a list of entities in a chunk position. If writing is not successful, an error is
	// returned.
	SaveEntities(position ChunkPos, entities []SaveableEntity) error
	// LoadBlockNBT loads the block NBT, also known as block entities, at a specific chunk position. If the
	// NBT cannot be read, LoadBlockNBT returns a non-nil error.
	LoadBlockNBT(position ChunkPos) ([]map[string]interface{}, error)
	// SaveBlockNBT saves block NBT, or block entities, to a specific chunk position. If the NBT cannot be
	// stored, SaveBlockNBT returns a non-nil error.
	SaveBlockNBT(position ChunkPos, data []map[string]interface{}) error
}

// NoIOProvider implements a Provider while not performing any disk I/O. It generates values on the run and
// dynamically, instead of reading and writing data, and returns otherwise empty values.
type NoIOProvider struct{}

// Settings ...
func (NoIOProvider) Settings() Settings { return defaultSettings() }

// SaveSettings ...
func (NoIOProvider) SaveSettings(Settings) {}

// LoadEntities ...
func (NoIOProvider) LoadEntities(ChunkPos) ([]SaveableEntity, error) {
	return nil, nil
}

// SaveEntities ...
func (NoIOProvider) SaveEntities(ChunkPos, []SaveableEntity) error {
	return nil
}

// LoadBlockNBT ...
func (NoIOProvider) LoadBlockNBT(ChunkPos) ([]map[string]interface{}, error) {
	return nil, nil
}

// SaveBlockNBT ...
func (NoIOProvider) SaveBlockNBT(ChunkPos, []map[string]interface{}) error {
	return nil
}

// SaveChunk ...
func (NoIOProvider) SaveChunk(ChunkPos, *chunk.Chunk) error {
	return nil
}

// LoadChunk ...
func (NoIOProvider) LoadChunk(ChunkPos) (*chunk.Chunk, bool, error) {
	return nil, false, nil
}

// Close ...
func (NoIOProvider) Close() error {
	return nil
}
