// (c) 2020-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package core

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/Toinounet21/swapeth/consensus/dummy"
	"github.com/Toinounet21/swapeth/core/rawdb"
	"github.com/Toinounet21/swapeth/core/state"
	"github.com/Toinounet21/swapeth/core/state/pruner"
	"github.com/Toinounet21/swapeth/core/types"
	"github.com/Toinounet21/swapeth/core/vm"
	"github.com/Toinounet21/swapeth/ethdb"
	"github.com/Toinounet21/swapeth/params"
	"github.com/ethereum/go-ethereum/common"
)

func TestArchiveBlockChain(t *testing.T) {
	create := func(db ethdb.Database, chainConfig *params.ChainConfig, lastAcceptedHash common.Hash) (*BlockChain, error) {
		// Import the chain. This runs all block validation rules.
		blockchain, err := NewBlockChain(
			db,
			&CacheConfig{
				TrieCleanLimit: 256,
				TrieDirtyLimit: 256,
				Pruning:        false, // Archive mode
				SnapshotLimit:  256,
			},
			chainConfig,
			dummy.NewDummyEngine(&dummy.ConsensusCallbacks{
				OnExtraStateChange: func(block *types.Block, sdb *state.StateDB) (*big.Int, *big.Int, error) {
					sdb.SetBalanceMultiCoin(common.HexToAddress("0xdeadbeef"), common.HexToHash("0xdeadbeef"), big.NewInt(block.Number().Int64()))
					return nil, nil, nil
				},
				OnFinalizeAndAssemble: func(header *types.Header, sdb *state.StateDB, txs []*types.Transaction) ([]byte, *big.Int, *big.Int, error) {
					sdb.SetBalanceMultiCoin(common.HexToAddress("0xdeadbeef"), common.HexToHash("0xdeadbeef"), big.NewInt(header.Number.Int64()))
					return nil, nil, nil, nil
				},
			}),
			vm.Config{},
			lastAcceptedHash,
		)
		return blockchain, err
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			tt.testFunc(t, create)
		})
	}
}

func TestArchiveBlockChainSnapsDisabled(t *testing.T) {
	create := func(db ethdb.Database, chainConfig *params.ChainConfig, lastAcceptedHash common.Hash) (*BlockChain, error) {
		// Import the chain. This runs all block validation rules.
		blockchain, err := NewBlockChain(
			db,
			&CacheConfig{
				TrieCleanLimit: 256,
				TrieDirtyLimit: 256,
				Pruning:        false, // Archive mode
				SnapshotLimit:  0,     // Disable snapshots
			},
			chainConfig,
			dummy.NewDummyEngine(&dummy.ConsensusCallbacks{
				OnExtraStateChange: func(block *types.Block, sdb *state.StateDB) (*big.Int, *big.Int, error) {
					sdb.SetBalanceMultiCoin(common.HexToAddress("0xdeadbeef"), common.HexToHash("0xdeadbeef"), big.NewInt(block.Number().Int64()))
					return nil, nil, nil
				},
				OnFinalizeAndAssemble: func(header *types.Header, sdb *state.StateDB, txs []*types.Transaction) ([]byte, *big.Int, *big.Int, error) {
					sdb.SetBalanceMultiCoin(common.HexToAddress("0xdeadbeef"), common.HexToHash("0xdeadbeef"), big.NewInt(header.Number.Int64()))
					return nil, nil, nil, nil
				},
			}),
			vm.Config{},
			lastAcceptedHash,
		)
		return blockchain, err
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			tt.testFunc(t, create)
		})
	}
}

func TestPruningBlockChain(t *testing.T) {
	create := func(db ethdb.Database, chainConfig *params.ChainConfig, lastAcceptedHash common.Hash) (*BlockChain, error) {
		// Import the chain. This runs all block validation rules.
		blockchain, err := NewBlockChain(
			db,
			&CacheConfig{
				TrieCleanLimit: 256,
				TrieDirtyLimit: 256,
				Pruning:        true, // Enable pruning
				SnapshotLimit:  256,
			},
			chainConfig,
			dummy.NewDummyEngine(&dummy.ConsensusCallbacks{
				OnExtraStateChange: func(block *types.Block, sdb *state.StateDB) (*big.Int, *big.Int, error) {
					sdb.SetBalanceMultiCoin(common.HexToAddress("0xdeadbeef"), common.HexToHash("0xdeadbeef"), big.NewInt(block.Number().Int64()))
					return nil, nil, nil
				},
				OnFinalizeAndAssemble: func(header *types.Header, sdb *state.StateDB, txs []*types.Transaction) ([]byte, *big.Int, *big.Int, error) {
					sdb.SetBalanceMultiCoin(common.HexToAddress("0xdeadbeef"), common.HexToHash("0xdeadbeef"), big.NewInt(header.Number.Int64()))
					return nil, nil, nil, nil
				},
			}),
			vm.Config{},
			lastAcceptedHash,
		)
		return blockchain, err
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			tt.testFunc(t, create)
		})
	}
}

func TestPruningBlockChainSnapsDisabled(t *testing.T) {
	create := func(db ethdb.Database, chainConfig *params.ChainConfig, lastAcceptedHash common.Hash) (*BlockChain, error) {
		// Import the chain. This runs all block validation rules.
		blockchain, err := NewBlockChain(
			db,
			&CacheConfig{
				TrieCleanLimit: 256,
				TrieDirtyLimit: 256,
				Pruning:        true, // Enable pruning
				SnapshotLimit:  0,    // Disable snapshots
			},
			chainConfig,
			dummy.NewDummyEngine(&dummy.ConsensusCallbacks{
				OnExtraStateChange: func(block *types.Block, sdb *state.StateDB) (*big.Int, *big.Int, error) {
					sdb.SetBalanceMultiCoin(common.HexToAddress("0xdeadbeef"), common.HexToHash("0xdeadbeef"), big.NewInt(block.Number().Int64()))
					return nil, nil, nil
				},
				OnFinalizeAndAssemble: func(header *types.Header, sdb *state.StateDB, txs []*types.Transaction) ([]byte, *big.Int, *big.Int, error) {
					sdb.SetBalanceMultiCoin(common.HexToAddress("0xdeadbeef"), common.HexToHash("0xdeadbeef"), big.NewInt(header.Number.Int64()))
					return nil, nil, nil, nil
				},
			}),
			vm.Config{},
			lastAcceptedHash,
		)
		return blockchain, err
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			tt.testFunc(t, create)
		})
	}
}

type wrappedStateManager struct {
	TrieWriter
}

func (w *wrappedStateManager) Shutdown() error { return nil }

func TestPruningBlockChainUngracefulShutdown(t *testing.T) {
	create := func(db ethdb.Database, chainConfig *params.ChainConfig, lastAcceptedHash common.Hash) (*BlockChain, error) {
		// Import the chain. This runs all block validation rules.
		blockchain, err := NewBlockChain(
			db,
			&CacheConfig{
				TrieCleanLimit: 256,
				TrieDirtyLimit: 256,
				Pruning:        true, // Enable pruning
				SnapshotLimit:  256,
			},
			chainConfig,
			dummy.NewDummyEngine(&dummy.ConsensusCallbacks{
				OnExtraStateChange: func(block *types.Block, sdb *state.StateDB) (*big.Int, *big.Int, error) {
					sdb.SetBalanceMultiCoin(common.HexToAddress("0xdeadbeef"), common.HexToHash("0xdeadbeef"), big.NewInt(block.Number().Int64()))
					return nil, nil, nil
				},
				OnFinalizeAndAssemble: func(header *types.Header, sdb *state.StateDB, txs []*types.Transaction) ([]byte, *big.Int, *big.Int, error) {
					sdb.SetBalanceMultiCoin(common.HexToAddress("0xdeadbeef"), common.HexToHash("0xdeadbeef"), big.NewInt(header.Number.Int64()))
					return nil, nil, nil, nil
				},
			}),
			vm.Config{},
			lastAcceptedHash,
		)
		if err != nil {
			return nil, err
		}

		// Overwrite state manager, so that Shutdown is not called.
		// This tests to ensure that the state manager handles an ungraceful shutdown correctly.
		blockchain.stateManager = &wrappedStateManager{TrieWriter: blockchain.stateManager}
		return blockchain, err
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			tt.testFunc(t, create)
		})
	}
}

func TestPruningBlockChainUngracefulShutdownSnapsDisabled(t *testing.T) {
	create := func(db ethdb.Database, chainConfig *params.ChainConfig, lastAcceptedHash common.Hash) (*BlockChain, error) {
		// Import the chain. This runs all block validation rules.
		blockchain, err := NewBlockChain(
			db,
			&CacheConfig{
				TrieCleanLimit: 256,
				TrieDirtyLimit: 256,
				Pruning:        true, // Enable pruning
				SnapshotLimit:  0,    // Disable snapshots
			},
			chainConfig,
			dummy.NewDummyEngine(&dummy.ConsensusCallbacks{
				OnExtraStateChange: func(block *types.Block, sdb *state.StateDB) (*big.Int, *big.Int, error) {
					sdb.SetBalanceMultiCoin(common.HexToAddress("0xdeadbeef"), common.HexToHash("0xdeadbeef"), big.NewInt(block.Number().Int64()))
					return nil, nil, nil
				},
				OnFinalizeAndAssemble: func(header *types.Header, sdb *state.StateDB, txs []*types.Transaction) ([]byte, *big.Int, *big.Int, error) {
					sdb.SetBalanceMultiCoin(common.HexToAddress("0xdeadbeef"), common.HexToHash("0xdeadbeef"), big.NewInt(header.Number.Int64()))
					return nil, nil, nil, nil
				},
			}),
			vm.Config{},
			lastAcceptedHash,
		)
		if err != nil {
			return nil, err
		}

		// Overwrite state manager, so that Shutdown is not called.
		// This tests to ensure that the state manager handles an ungraceful shutdown correctly.
		blockchain.stateManager = &wrappedStateManager{TrieWriter: blockchain.stateManager}
		return blockchain, err
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			tt.testFunc(t, create)
		})
	}
}

func TestEnableSnapshots(t *testing.T) {
	// Set snapshots to be disabled the first time, and then enable them on the restart
	snapLimit := 0
	create := func(db ethdb.Database, chainConfig *params.ChainConfig, lastAcceptedHash common.Hash) (*BlockChain, error) {
		// Import the chain. This runs all block validation rules.
		blockchain, err := NewBlockChain(
			db,
			&CacheConfig{
				TrieCleanLimit: 256,
				TrieDirtyLimit: 256,
				Pruning:        true,      // Enable pruning
				SnapshotLimit:  snapLimit, // Disable snapshots
			},
			chainConfig,
			dummy.NewDummyEngine(&dummy.ConsensusCallbacks{
				OnExtraStateChange: func(block *types.Block, sdb *state.StateDB) (*big.Int, *big.Int, error) {
					sdb.SetBalanceMultiCoin(common.HexToAddress("0xdeadbeef"), common.HexToHash("0xdeadbeef"), big.NewInt(block.Number().Int64()))
					return nil, nil, nil
				},
				OnFinalizeAndAssemble: func(header *types.Header, sdb *state.StateDB, txs []*types.Transaction) ([]byte, *big.Int, *big.Int, error) {
					sdb.SetBalanceMultiCoin(common.HexToAddress("0xdeadbeef"), common.HexToHash("0xdeadbeef"), big.NewInt(header.Number.Int64()))
					return nil, nil, nil, nil
				},
			}),
			vm.Config{},
			lastAcceptedHash,
		)
		if err != nil {
			return nil, err
		}
		snapLimit = 256

		return blockchain, err
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			tt.testFunc(t, create)
		})
	}
}

func TestCorruptSnapshots(t *testing.T) {
	create := func(db ethdb.Database, chainConfig *params.ChainConfig, lastAcceptedHash common.Hash) (*BlockChain, error) {
		// Delete the snapshot block hash and state root to ensure that if we die in between writing a snapshot
		// diff layer to disk at any point, we can still recover on restart.
		rawdb.DeleteSnapshotBlockHash(db)
		rawdb.DeleteSnapshotRoot(db)
		// Import the chain. This runs all block validation rules.
		blockchain, err := NewBlockChain(
			db,
			&CacheConfig{
				TrieCleanLimit: 256,
				TrieDirtyLimit: 256,
				Pruning:        true, // Enable pruning
				SnapshotLimit:  256,  // Disable snapshots
			},
			chainConfig,
			dummy.NewDummyEngine(&dummy.ConsensusCallbacks{
				OnExtraStateChange: func(block *types.Block, sdb *state.StateDB) (*big.Int, *big.Int, error) {
					sdb.SetBalanceMultiCoin(common.HexToAddress("0xdeadbeef"), common.HexToHash("0xdeadbeef"), big.NewInt(block.Number().Int64()))
					return nil, nil, nil
				},
				OnFinalizeAndAssemble: func(header *types.Header, sdb *state.StateDB, txs []*types.Transaction) ([]byte, *big.Int, *big.Int, error) {
					sdb.SetBalanceMultiCoin(common.HexToAddress("0xdeadbeef"), common.HexToHash("0xdeadbeef"), big.NewInt(header.Number.Int64()))
					return nil, nil, nil, nil
				},
			}),
			vm.Config{},
			lastAcceptedHash,
		)
		if err != nil {
			return nil, err
		}

		return blockchain, err
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			tt.testFunc(t, create)
		})
	}
}

func TestBlockChainOfflinePruningUngracefulShutdown(t *testing.T) {
	create := func(db ethdb.Database, chainConfig *params.ChainConfig, lastAcceptedHash common.Hash) (*BlockChain, error) {
		// Import the chain. This runs all block validation rules.
		blockchain, err := NewBlockChain(
			db,
			&CacheConfig{
				TrieCleanLimit: 256,
				TrieDirtyLimit: 256,
				Pruning:        true, // Enable pruning
				SnapshotLimit:  256,
			},
			chainConfig,
			dummy.NewDummyEngine(&dummy.ConsensusCallbacks{
				OnExtraStateChange: func(block *types.Block, sdb *state.StateDB) (*big.Int, *big.Int, error) {
					sdb.SetBalanceMultiCoin(common.HexToAddress("0xdeadbeef"), common.HexToHash("0xdeadbeef"), big.NewInt(block.Number().Int64()))
					return nil, nil, nil
				},
				OnFinalizeAndAssemble: func(header *types.Header, sdb *state.StateDB, txs []*types.Transaction) ([]byte, *big.Int, *big.Int, error) {
					sdb.SetBalanceMultiCoin(common.HexToAddress("0xdeadbeef"), common.HexToHash("0xdeadbeef"), big.NewInt(header.Number.Int64()))
					return nil, nil, nil, nil
				},
			}),
			vm.Config{},
			lastAcceptedHash,
		)
		if err != nil {
			return nil, err
		}

		// Overwrite state manager, so that Shutdown is not called.
		// This tests to ensure that the state manager handles an ungraceful shutdown correctly.
		blockchain.stateManager = &wrappedStateManager{TrieWriter: blockchain.stateManager}

		if lastAcceptedHash == (common.Hash{}) {
			return blockchain, nil
		}

		tempDir := t.TempDir()
		if err := blockchain.CleanBlockRootsAboveLastAccepted(); err != nil {
			return nil, err
		}
		pruner, err := pruner.NewPruner(db, tempDir, 256)
		if err != nil {
			return nil, fmt.Errorf("offline pruning failed (%s, %d): %w", tempDir, 256, err)
		}

		targetRoot := blockchain.LastAcceptedBlock().Root()
		if err := pruner.Prune(targetRoot); err != nil {
			return nil, fmt.Errorf("failed to prune blockchain with target root: %s due to: %w", targetRoot, err)
		}
		// Re-initialize the blockchain after pruning
		return NewBlockChain(
			db,
			&CacheConfig{
				TrieCleanLimit: 256,
				TrieDirtyLimit: 256,
				Pruning:        true, // Enable pruning
				SnapshotLimit:  256,
			},
			chainConfig,
			dummy.NewDummyEngine(&dummy.ConsensusCallbacks{
				OnExtraStateChange: func(block *types.Block, sdb *state.StateDB) (*big.Int, *big.Int, error) {
					sdb.SetBalanceMultiCoin(common.HexToAddress("0xdeadbeef"), common.HexToHash("0xdeadbeef"), big.NewInt(block.Number().Int64()))
					return nil, nil, nil
				},
				OnFinalizeAndAssemble: func(header *types.Header, sdb *state.StateDB, txs []*types.Transaction) ([]byte, *big.Int, *big.Int, error) {
					sdb.SetBalanceMultiCoin(common.HexToAddress("0xdeadbeef"), common.HexToHash("0xdeadbeef"), big.NewInt(header.Number.Int64()))
					return nil, nil, nil, nil
				},
			}),
			vm.Config{},
			lastAcceptedHash,
		)
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			tt.testFunc(t, create)
		})
	}
}
