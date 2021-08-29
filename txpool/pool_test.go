/*
   Copyright 2021 Erigon contributors

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package txpool

import (
	"testing"

	"github.com/google/btree"
	"github.com/holiman/uint256"
	"github.com/ledgerwatch/erigon-lib/kv/memdb"
	"github.com/stretchr/testify/require"
)

func TestSenders(t *testing.T) {
	t.Run("evict_all_on_next_round", func(t *testing.T) {
		senders, require := NewSendersCache(), require.New(t)
		senders.senderInfo[1] = newSenderInfo(1, *uint256.NewInt(1))
		senders.senderInfo[2] = newSenderInfo(1, *uint256.NewInt(1))
		_, tx := memdb.NewTestPoolTx(t)
		byNonce := &ByNonce{btree.New(16)}

		evicted, err := senders.flush(tx, byNonce, []uint64{1, 2}, 1)
		require.NoError(err)
		require.Zero(evicted)
		evicted, err = senders.flush(tx, byNonce, []uint64{}, 1)
		require.NoError(err)
		require.Equal(2, int(evicted))
	})
	t.Run("do_not_evict_if_used_in_current_round", func(t *testing.T) {
		senders, require := NewSendersCache(), require.New(t)
		_, tx := memdb.NewTestPoolTx(t)
		byNonce := &ByNonce{btree.New(16)}

		senders.senderInfo[1] = newSenderInfo(1, *uint256.NewInt(1))
		senders.senderInfo[2] = newSenderInfo(1, *uint256.NewInt(1))
		evicted, err := senders.flush(tx, byNonce, []uint64{1, 2}, 1)
		require.NoError(err)
		require.Zero(evicted)
		senders.senderInfo[1] = newSenderInfo(1, *uint256.NewInt(1)) // means used in current round, but still has 0 transactions
		evicted, err = senders.flush(tx, byNonce, []uint64{1}, 1)
		require.NoError(err)
		require.Equal(1, int(evicted))
	})

}
