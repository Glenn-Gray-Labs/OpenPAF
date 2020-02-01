package core

import (
	"fmt"
	math32 "github.com/chewxy/math32"
	"math/rand"
)

// VoxelGrid is a 3-dimensional grid of voxels based on a given set of dimensions and a resolution per
// dimensional unit. The purpose of this type is only to manage the contents of a collection of voxels.
type VoxelGrid struct {
	resolution uint
	counts     [3]uint // dimensions in number of voxels
	voxels     []Voxel // the stored voxels for the grid
}

// NewVoxelGrid calculates the requirements for a voxel grid, configures one, and then returns the result.
func NewVoxelGrid(width, height, depth float32, resolution uint) *VoxelGrid {
	w := uint(math32.Ceil(math32.Abs(width * float32(resolution))))
	h := uint(math32.Ceil(math32.Abs(height * float32(resolution))))
	d := uint(math32.Ceil(math32.Abs(depth * float32(resolution))))

	return &VoxelGrid{
		resolution: resolution,
		counts:     [3]uint{w, h, d},
		voxels:     make([]Voxel, w*h*d),
	}
}

// getIndex returns the slice index based on a given set of 3D coordinates.
func (vg *VoxelGrid) getIndex(x, y, z uint) uint {
	w, h, d := vg.counts[0], vg.counts[1], vg.counts[2]

	if x >= w || y >= h || z >= d {
		panic(fmt.Errorf("x, y, z coordinate of %d, %d, %d exceed max voxel dimensions of %d, %d, %d", x, y, z, w, h, d))
	}

	return (w * h * z) + (w * y) + x
}

// getCoordinate converts a slice index for the grid to the equivalent 3D coordinates.
func (vg *VoxelGrid) getCoordinate(index uint) (x, y, z uint) {
	if index > uint(len(vg.voxels)) {
		panic(fmt.Errorf("given index of %d exceeds the size of the voxel grid", index))
	}

	zquo, z := divmod(index, vg.counts[2])
	yquo, y := divmod(zquo, vg.counts[1])
	x = yquo % vg.counts[0]
	return
}

// Get returns the value of the voxel stored at the given 3D coordinates.
func (vg *VoxelGrid) Get(x, y, z uint) Voxel {
	return vg.voxels[vg.getIndex(x, y, z)]
}

// Set will set the value of a voxel using the 3D coordinates.
func (vg *VoxelGrid) Set(x, y, z uint, vox Voxel) {
	vg.voxels[vg.getIndex(x, y, z)] = vox
}

// Copy creates a new voxel grid instance with the some configuration and contents of this one.
func (vg *VoxelGrid) Copy() *VoxelGrid {
	newGrid := &VoxelGrid{counts: vg.counts}
	copy(newGrid.voxels, vg.voxels)
	return newGrid
}

// Fill sets all of the voxels in the grid to same the voxel value.
func (vg *VoxelGrid) Fill(value float32) *VoxelGrid {
	value = math32.Min(1, math32.Max(0, value))
	for i := range vg.voxels {
		vg.voxels[i] = Voxel(value)
	}

	return vg
}

// Randomize generates a random value for each voxel in the grid. This is largely for testing purposes and
// will eventually be implemented as a graph procedure.
func (vg *VoxelGrid) Randomize(seed int64) *VoxelGrid {
	r := rand.New(rand.NewSource(seed))
	for i := range vg.voxels {
		vg.voxels[i] = Voxel(r.Float32())
	}

	return vg
}

// HighPass runs over each voxel in the grid and then sets the voxel's value to either 0 or 1 depending on
// whether or not its within the given range of tolerance. This will eventually be implemented as a graph procedure.
func (vg *VoxelGrid) HighPass(tolerance float32) *VoxelGrid {
	for i, v := range vg.voxels {
		vg.voxels[i] = 0
		if float32(v) > tolerance {
			vg.voxels[i] = 1
		}
	}

	return vg
}

// VertexPoints converts all voxels with a value greater than 0.5 to a list of 3D vertex
// coordinates.
func (vg *VoxelGrid) VertexPoints() [][3]float32 {
	var points [][3]float32
	for i, v := range vg.voxels {
		if v < 0.5 {
			continue
		}

		x, y, z := vg.getCoordinate(uint(i))
		res := float32(vg.resolution)
		points = append(points, [3]float32{
			float32(x) / res,
			float32(y) / res,
			float32(z) / res,
		})
	}

	return points
}

// Mesh attempts to perform cube march algorithm using the current voxel grid and returns the resulting
// mesh information with the results.
func (vg *VoxelGrid) Mesh() {
	// TODO: implement this function
}
