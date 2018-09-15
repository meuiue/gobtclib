package client

import (
	"encoding/json"
	"github.com/chainlibs/gobtclib/utils"
	"bytes"
	"encoding/hex"
)

/*
Description:
FutureGetBestBlockHashResult is a future promise to deliver the result of a
GetBestBlockAsync RPC invocation (or an applicable error).
 * Author: architect.bian
 * Date: 2018/09/14 13:20
 */
type FutureGetBestBlockHashResult chan *response

/*
Description:
Receive waits for the response promised by the future and returns the hash of
the best block in the longest block chain.
 * Author: architect.bian
 * Date: 2018/09/14 13:20
 */
func (r FutureGetBestBlockHashResult) Receive() (*Hash, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal result as a string.
	var txHashStr string
	err = json.Unmarshal(res, &txHashStr)
	if err != nil {
		return nil, err
	}
	return NewHashFromStr(txHashStr)
}
/*
Description:
GetBestBlockHashCmd defines the getbestblockhash JSON-RPC command.
 * Author: architect.bian
 * Date: 2018/09/14 13:19
 */
type GetBestBlockHashCmd struct{}

/*
Description:
GetBestBlockHashAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See GetBestBlockHash for the blocking version and more details.
 * Author: architect.bian
 * Date: 2018/09/14 13:18
 */
func (c *Client) GetBestBlockHashAsync() FutureGetBestBlockHashResult {
	cmd := &GetBestBlockHashCmd{}
	return c.sendCmd(cmd)
}

/*
Description:
GetBestBlockHash returns the hash of the best block in the longest block chain.
 * Author: architect.bian
 * Date: 2018/09/14 13:15
 */
func (c *Client) GetBestBlockHash() (*Hash, error) {
	return c.GetBestBlockHashAsync().Receive()
}

/*
Description:
FutureGetBlockCountResult is a future promise to deliver the result of a
GetBlockCountAsync RPC invocation (or an applicable error).
 * Author: architect.bian
 * Date: 2018/09/14 12:55
 */
type FutureGetBlockCountResult chan *response

/*
Description:
Receive waits for the response promised by the future and returns the number
of blocks in the longest block chain.
 * Author: architect.bian
 * Date: 2018/09/14 12:59
 */
func (r FutureGetBlockCountResult) Receive() (int64, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return -1, err
	}

	// Unmarshal the result as an int64.
	var count int64
	err = json.Unmarshal(res, &count)
	if err != nil {
		return -1, err
	}
	return count, nil
}

/*
Description:
GetBlockCountCmd defines the getblockcount JSON-RPC command.
 * Author: architect.bian
 * Date: 2018/09/14 12:50
 */
type GetBlockCountCmd struct{}

/*
Description:
GetBlockCountAsync returns an instance of a type that can be used to get the
result of the RPC at some future time by invoking the Receive function on the
returned instance.

See GetBlockCount for the blocking version and more details.
 * Author: architect.bian
 * Date: 2018/09/14 12:49
 */
func (c *Client) GetBlockCountAsync() FutureGetBlockCountResult {
	cmd := &GetBlockCountCmd{}
	return c.sendCmd(cmd)
}

/*
Description:
GetBlockCount returns the number of blocks in the longest block chain.
 * Author: architect.bian
 * Date: 2018/09/14 12:46
 */
func (c *Client) GetBlockCount() (int64, error) {
	return c.GetBlockCountAsync().Receive()
}

/*
Description:
FutureGetDifficultyResult is a future promise to deliver the result of a
GetDifficultyAsync RPC invocation (or an applicable error).
 * Author: architect.bian
 * Date: 2018/09/14 17:36
 */
type FutureGetDifficultyResult chan *response

/*
Description:
Receive waits for the response promised by the future and returns the
proof-of-work difficulty as a multiple of the minimum difficulty.
 * Author: architect.bian
 * Date: 2018/09/14 17:36
 */
func (r FutureGetDifficultyResult) Receive() (float64, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return 0, err
	}

	// Unmarshal the result as a float64.
	var difficulty float64
	err = json.Unmarshal(res, &difficulty)
	if err != nil {
		return 0, err
	}
	return difficulty, nil
}

/*
Description:
GetDifficultyCmd defines the getdifficulty JSON-RPC command.
 * Author: architect.bian
 * Date: 2018/09/14 17:36
 */
type GetDifficultyCmd struct{}

/*
Description:
GetDifficultyAsync returns an instance of a type that can be used to get the
result of the RPC at some future time by invoking the Receive function on the
returned instance.

See GetDifficulty for the blocking version and more details.
 * Author: architect.bian
 * Date: 2018/09/14 17:35
 */
func (c *Client) GetDifficultyAsync() FutureGetDifficultyResult {
	cmd := &GetDifficultyCmd{}
	return c.sendCmd(cmd)
}

/*
Description:
GetDifficulty returns the proof-of-work difficulty as a multiple of the
minimum difficulty. The result is bits/params.PowLimitBits
 * Author: architect.bian
 * Date: 2018/09/14 17:34
 */
func (c *Client) GetDifficulty() (float64, error) {
	return c.GetDifficultyAsync().Receive()
}

/*
Description:
FutureGetBlockHashResult is a future promise to deliver the result of a
GetBlockHashAsync RPC invocation (or an applicable error).
 * Author: architect.bian
 * Date: 2018/09/15 15:44
 */
type FutureGetBlockHashResult chan *response

/*
Description:
Receive waits for the response promised by the future and returns the hash of
the block in the best block chain at the given height.
 * Author: architect.bian
 * Date: 2018/09/15 15:45
 */
func (r FutureGetBlockHashResult) Receive() (*Hash, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal the result as a string-encoded sha.
	var txHashStr string
	err = json.Unmarshal(res, &txHashStr)
	if err != nil {
		return nil, err
	}
	return NewHashFromStr(txHashStr)
}

/*
Description:
GetBlockHashCmd defines the getblockhash JSON-RPC command.
 * Author: architect.bian
 * Date: 2018/09/15 15:44
 */
type GetBlockHashCmd struct {
	Index int64
}

/*
Description:
GetBlockHashAsync returns an instance of a type that can be used to get the
result of the RPC at some future time by invoking the Receive function on the
returned instance.

See GetBlockHash for the blocking version and more details.
 * Author: architect.bian
 * Date: 2018/09/15 15:35
 */
func (c *Client) GetBlockHashAsync(blockHeight int64) FutureGetBlockHashResult {
	cmd := &GetBlockHashCmd{
		Index: blockHeight,
	}
	return c.sendCmd(cmd)
}

/*
Description:
GetBlockHash returns the hash of the block in the best block chain at the given height.
 * Author: architect.bian
 * Date: 2018/09/15 15:34
 */
func (c *Client) GetBlockHash(blockHeight int64) (*Hash, error) {
	return c.GetBlockHashAsync(blockHeight).Receive()
}

/*
Description:
FutureGetBlockHeaderVerboseResult is a future promise to deliver the result of a
GetBlockAsync RPC invocation (or an applicable error).
 * Author: architect.bian
 * Date: 2018/09/15 17:52
 */
type FutureGetBlockHeaderVerboseResult chan *response

/*
Description:
Receive waits for the response promised by the future and returns the
data structure of the blockheader requested from the server given its hash.
 * Author: architect.bian
 * Date: 2018/09/15 17:52
 */
func (r FutureGetBlockHeaderVerboseResult) Receive() (*GetBlockHeaderVerboseResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal result as a string.
	var bh GetBlockHeaderVerboseResult
	err = json.Unmarshal(res, &bh)
	if err != nil {
		return nil, err
	}

	return &bh, nil
}

/*
Description:
GetBlockHeaderCmd defines the getblockheader JSON-RPC command.
 * Author: architect.bian
 * Date: 2018/09/15 17:52
 */
type GetBlockHeaderCmd struct {
	Hash    string
	Verbose *bool `jsonrpcdefault:"true"`
}

/*
Description:
GetBlockHeaderVerboseAsync returns an instance of a type that can be used to get the
result of the RPC at some future time by invoking the Receive function on the
returned instance.

See GetBlockHeader for the blocking version and more details.
 * Author: architect.bian
 * Date: 2018/09/15 17:51
 */
func (c *Client) GetBlockHeaderVerboseAsync(blockHash *Hash) FutureGetBlockHeaderVerboseResult {
	hash := ""
	if blockHash != nil {
		hash = blockHash.String()
	}

	cmd := &GetBlockHeaderCmd{
		Hash:    hash,
		Verbose: utils.Basis.Bool(true),
	}
	return c.sendCmd(cmd)
}

/*
Description:
GetBlockHeaderVerbose returns a data structure with information about the
blockheader from the server given its hash.

See GetBlockHeader to retrieve a blockheader instead.
 * Author: architect.bian
 * Date: 2018/09/15 17:51
 */
func (c *Client) GetBlockHeaderVerbose(blockHash *Hash) (*GetBlockHeaderVerboseResult, error) {
	return c.GetBlockHeaderVerboseAsync(blockHash).Receive()
}

// FutureGetMempoolEntryResult is a future promise to deliver the result of a
// GetMempoolEntryAsync RPC invocation (or an applicable error).
type FutureGetMempoolEntryResult chan *response

// Receive waits for the response promised by the future and returns a data
// structure with information about the transaction in the memory pool given
// its hash.
func (r FutureGetMempoolEntryResult) Receive() (*GetMempoolEntryResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal the result as an array of strings.
	var mempoolEntryResult GetMempoolEntryResult
	err = json.Unmarshal(res, &mempoolEntryResult)
	if err != nil {
		return nil, err
	}

	return &mempoolEntryResult, nil
}

// GetMempoolEntryCmd defines the getmempoolentry JSON-RPC command.
type GetMempoolEntryCmd struct {
	TxID string
}

// GetMempoolEntryAsync returns an instance of a type that can be used to get the
// result of the RPC at some future time by invoking the Receive function on the
// returned instance.
//
// See GetMempoolEntry for the blocking version and more details.
func (c *Client) GetMempoolEntryAsync(txHash string) FutureGetMempoolEntryResult {
	cmd := &GetMempoolEntryCmd{
		TxID: txHash,
	}
	return c.sendCmd(cmd)
}

// GetMempoolEntry returns a data structure with information about the
// transaction in the memory pool given its hash.
func (c *Client) GetMempoolEntry(txHash string) (*GetMempoolEntryResult, error) {
	return c.GetMempoolEntryAsync(txHash).Receive()
}

// FutureGetRawMempoolResult is a future promise to deliver the result of a
// GetRawMempoolAsync RPC invocation (or an applicable error).
type FutureGetRawMempoolResult chan *response

// Receive waits for the response promised by the future and returns the hashes
// of all transactions in the memory pool.
func (r FutureGetRawMempoolResult) Receive() ([]*Hash, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal the result as an array of strings.
	var txHashStrs []string
	err = json.Unmarshal(res, &txHashStrs)
	if err != nil {
		return nil, err
	}

	// Create a slice of ShaHash arrays from the string slice.
	txHashes := make([]*Hash, 0, len(txHashStrs))
	for _, hashStr := range txHashStrs {
		txHash, err := NewHashFromStr(hashStr)
		if err != nil {
			return nil, err
		}
		txHashes = append(txHashes, txHash)
	}

	return txHashes, nil
}

// GetRawMempoolCmd defines the getmempool JSON-RPC command.
type GetRawMempoolCmd struct {
	Verbose *bool `jsonrpcdefault:"false"`
}

// GetRawMempoolAsync returns an instance of a type that can be used to get the
// result of the RPC at some future time by invoking the Receive function on the
// returned instance.
//
// See GetRawMempool for the blocking version and more details.
func (c *Client) GetRawMempoolAsync() FutureGetRawMempoolResult {
	cmd := &GetRawMempoolCmd{
		Verbose: utils.Basis.Bool(false),
	}
	return c.sendCmd(cmd)
}

// GetRawMempool returns the hashes of all transactions in the memory pool.
//
// See GetRawMempoolVerbose to retrieve data structures with information about
// the transactions instead.
func (c *Client) GetRawMempool() ([]*Hash, error) {
	return c.GetRawMempoolAsync().Receive()
}

// FutureInvalidateBlockResult is a future promise to deliver the result of a
// InvalidateBlockAsync RPC invocation (or an applicable error).
type FutureInvalidateBlockResult chan *response

// Receive waits for the response promised by the future and returns the raw
// block requested from the server given its hash.
func (r FutureInvalidateBlockResult) Receive() error {
	_, err := receiveFuture(r)

	return err
}

// InvalidateBlockCmd defines the invalidateblock JSON-RPC command.
type InvalidateBlockCmd struct {
	BlockHash string
}

// InvalidateBlockAsync returns an instance of a type that can be used to get the
// result of the RPC at some future time by invoking the Receive function on the
// returned instance.
//
// See InvalidateBlock for the blocking version and more details.
func (c *Client) InvalidateBlockAsync(blockHash *Hash) FutureInvalidateBlockResult {
	hash := ""
	if blockHash != nil {
		hash = blockHash.String()
	}

	cmd := &InvalidateBlockCmd{
		BlockHash: hash,
	}
	return c.sendCmd(cmd)
}

// InvalidateBlock invalidates a specific block.
func (c *Client) InvalidateBlock(blockHash *Hash) error {
	return c.InvalidateBlockAsync(blockHash).Receive()
}

// FutureGetBlockChainInfoResult is a promise to deliver the result of a
// GetBlockChainInfoAsync RPC invocation (or an applicable error).
type FutureGetBlockChainInfoResult chan *response

// Receive waits for the response promised by the future and returns chain info
// result provided by the server.
func (r FutureGetBlockChainInfoResult) Receive() (*GetBlockChainInfoResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	var chainInfo GetBlockChainInfoResult
	if err := json.Unmarshal(res, &chainInfo); err != nil {
		return nil, err
	}
	return &chainInfo, nil
}
// GetBlockChainInfoCmd defines the getblockchaininfo JSON-RPC command.
type GetBlockChainInfoCmd struct{}

// GetBlockChainInfoAsync returns an instance of a type that can be used to get
// the result of the RPC at some future time by invoking the Receive function
// on the returned instance.
//
// See GetBlockChainInfo for the blocking version and more details.
func (c *Client) GetBlockChainInfoAsync() FutureGetBlockChainInfoResult {
	cmd := &GetBlockChainInfoCmd{}
	return c.sendCmd(cmd)
}

// GetBlockChainInfo returns information related to the processing state of
// various chain-specific details such as the current difficulty from the tip
// of the main chain.
func (c *Client) GetBlockChainInfo() (*GetBlockChainInfoResult, error) {
	return c.GetBlockChainInfoAsync().Receive()
}

// FutureGetBlockHeaderResult is a future promise to deliver the result of a
// GetBlockHeaderAsync RPC invocation (or an applicable error).
type FutureGetBlockHeaderResult chan *response

// Receive waits for the response promised by the future and returns the
// blockheader requested from the server given its hash.
func (r FutureGetBlockHeaderResult) Receive() (*BlockHeader, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal result as a string.
	var bhHex string
	err = json.Unmarshal(res, &bhHex)
	if err != nil {
		return nil, err
	}

	serializedBH, err := hex.DecodeString(bhHex)
	if err != nil {
		return nil, err
	}

	// Deserialize the blockheader and return it.
	var bh BlockHeader
	err = bh.Deserialize(bytes.NewReader(serializedBH))
	if err != nil {
		return nil, err
	}

	return &bh, err
}

// GetBlockHeaderAsync returns an instance of a type that can be used to get the
// result of the RPC at some future time by invoking the Receive function on the
// returned instance.
//
// See GetBlockHeader for the blocking version and more details.
func (c *Client) GetBlockHeaderAsync(blockHash *Hash) FutureGetBlockHeaderResult {
	hash := ""
	if blockHash != nil {
		hash = blockHash.String()
	}

	cmd := &GetBlockHeaderCmd{
		Hash:    hash,
		Verbose: utils.Basis.Bool(false),
	}
	return c.sendCmd(cmd)
}

// GetBlockHeader returns the blockheader from the server given its hash.
//
// See GetBlockHeaderVerbose to retrieve a data structure with information about the
// block instead.
func (c *Client) GetBlockHeader(blockHash *Hash) (*BlockHeader, error) {
	return c.GetBlockHeaderAsync(blockHash).Receive()
}

// FutureGetBlockVerboseResult is a future promise to deliver the result of a
// GetBlockVerboseAsync RPC invocation (or an applicable error).
type FutureGetBlockVerboseResult chan *response

// Receive waits for the response promised by the future and returns the data
// structure from the server with information about the requested block.
func (r FutureGetBlockVerboseResult) Receive() (*GetBlockVerboseResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal the raw result into a BlockResult.
	var blockResult GetBlockVerboseResult
	err = json.Unmarshal(res, &blockResult)
	if err != nil {
		return nil, err
	}
	return &blockResult, nil
}

// GetBlockCmd defines the getblock JSON-RPC command.
type GetBlockCmd struct {
	Hash      string
	Verbose   *bool `jsonrpcdefault:"true"`
	VerboseTx *bool `jsonrpcdefault:"false"`
}

// GetBlockVerboseAsync returns an instance of a type that can be used to get
// the result of the RPC at some future time by invoking the Receive function on
// the returned instance.
//
// See GetBlockVerbose for the blocking version and more details.
func (c *Client) GetBlockVerboseAsync(blockHash *Hash) FutureGetBlockVerboseResult {
	hash := ""
	if blockHash != nil {
		hash = blockHash.String()
	}

	cmd := &GetBlockCmd{
		Hash:      hash,
		Verbose:   utils.Basis.Bool(true),
		VerboseTx: nil,
	}
	return c.sendCmd(cmd)
}

// GetBlockVerbose returns a data structure from the server with information
// about a block given its hash.
//
// See GetBlockVerboseTx to retrieve transaction data structures as well.
// See GetBlock to retrieve a raw block instead.
func (c *Client) GetBlockVerbose(blockHash *Hash) (*GetBlockVerboseResult, error) {
	return c.GetBlockVerboseAsync(blockHash).Receive()
}

// FutureGetRawMempoolVerboseResult is a future promise to deliver the result of
// a GetRawMempoolVerboseAsync RPC invocation (or an applicable error).
type FutureGetRawMempoolVerboseResult chan *response

// Receive waits for the response promised by the future and returns a map of
// transaction hashes to an associated data structure with information about the
// transaction for all transactions in the memory pool.
func (r FutureGetRawMempoolVerboseResult) Receive() (map[string]GetRawMempoolVerboseResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal the result as a map of strings (tx shas) to their detailed
	// results.
	var mempoolItems map[string]GetRawMempoolVerboseResult
	err = json.Unmarshal(res, &mempoolItems)
	if err != nil {
		return nil, err
	}
	return mempoolItems, nil
}

// GetRawMempoolVerboseAsync returns an instance of a type that can be used to
// get the result of the RPC at some future time by invoking the Receive
// function on the returned instance.
//
// See GetRawMempoolVerbose for the blocking version and more details.
func (c *Client) GetRawMempoolVerboseAsync() FutureGetRawMempoolVerboseResult {
	cmd := &GetRawMempoolCmd{
		Verbose: utils.Basis.Bool(true),
	}
	return c.sendCmd(cmd)
}

// GetRawMempoolVerbose returns a map of transaction hashes to an associated
// data structure with information about the transaction for all transactions in
// the memory pool.
//
// See GetRawMempool to retrieve only the transaction hashes instead.
func (c *Client) GetRawMempoolVerbose() (map[string]GetRawMempoolVerboseResult, error) {
	return c.GetRawMempoolVerboseAsync().Receive()
}

// FutureEstimateFeeResult is a future promise to deliver the result of a
// EstimateFeeAsync RPC invocation (or an applicable error).
type FutureEstimateFeeResult chan *response

// Receive waits for the response promised by the future and returns the info
// provided by the server.
func (r FutureEstimateFeeResult) Receive() (float64, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return -1, err
	}

	// Unmarshal result as a getinfo result object.
	var fee float64
	err = json.Unmarshal(res, &fee)
	if err != nil {
		return -1, err
	}

	return fee, nil
}

// EstimateFeeAsync returns an instance of a type that can be used to get the result
// of the RPC at some future time by invoking the Receive function on the
// returned instance.
//
// See EstimateFee for the blocking version and more details.
//func (c *Client) EstimateFeeAsync(numBlocks int64) FutureEstimateFeeResult {
//	cmd := NewEstimateFeeCmd(numBlocks)
//	return c.sendCmd(cmd)
//}

// EstimateFee provides an estimated fee  in bitcoins per kilobyte.
//func (c *Client) EstimateFee(numBlocks int64) (float64, error) {
//	return c.EstimateFeeAsync(numBlocks).Receive()
//}

// FutureVerifyChainResult is a future promise to deliver the result of a
// VerifyChainAsync, VerifyChainLevelAsyncRPC, or VerifyChainBlocksAsync
// invocation (or an applicable error).
type FutureVerifyChainResult chan *response

// Receive waits for the response promised by the future and returns whether
// or not the chain verified based on the check level and number of blocks
// to verify specified in the original call.
func (r FutureVerifyChainResult) Receive() (bool, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return false, err
	}

	// Unmarshal the result as a boolean.
	var verified bool
	err = json.Unmarshal(res, &verified)
	if err != nil {
		return false, err
	}
	return verified, nil
}

// VerifyChainCmd defines the verifychain JSON-RPC command.
type VerifyChainCmd struct {
	CheckLevel *int32 `jsonrpcdefault:"3"`
	CheckDepth *int32 `jsonrpcdefault:"288"` // 0 = all
}

// VerifyChainAsync returns an instance of a type that can be used to get the
// result of the RPC at some future time by invoking the Receive function on the
// returned instance.
//
// See VerifyChain for the blocking version and more details.
func (c *Client) VerifyChainAsync() FutureVerifyChainResult {
	cmd := &VerifyChainCmd{
		CheckLevel: nil,
		CheckDepth: nil,
	}
	return c.sendCmd(cmd)
}

// VerifyChain requests the server to verify the block chain database using
// the default check level and number of blocks to verify.
//
// See VerifyChainLevel and VerifyChainBlocks to override the defaults.
func (c *Client) VerifyChain() (bool, error) {
	return c.VerifyChainAsync().Receive()
}

// VerifyChainLevelAsync returns an instance of a type that can be used to get
// the result of the RPC at some future time by invoking the Receive function on
// the returned instance.
//
// See VerifyChainLevel for the blocking version and more details.
func (c *Client) VerifyChainLevelAsync(checkLevel int32) FutureVerifyChainResult {
	cmd := &VerifyChainCmd{
		CheckLevel: &checkLevel,
		CheckDepth: nil,
	}
	return c.sendCmd(cmd)
}

// VerifyChainLevel requests the server to verify the block chain database using
// the passed check level and default number of blocks to verify.
//
// The check level controls how thorough the verification is with higher numbers
// increasing the amount of checks done as consequently how long the
// verification takes.
//
// See VerifyChain to use the default check level and VerifyChainBlocks to
// override the number of blocks to verify.
func (c *Client) VerifyChainLevel(checkLevel int32) (bool, error) {
	return c.VerifyChainLevelAsync(checkLevel).Receive()
}

// VerifyChainBlocksAsync returns an instance of a type that can be used to get
// the result of the RPC at some future time by invoking the Receive function on
// the returned instance.
//
// See VerifyChainBlocks for the blocking version and more details.
func (c *Client) VerifyChainBlocksAsync(checkLevel, numBlocks int32) FutureVerifyChainResult {
	cmd := &VerifyChainCmd{
		CheckLevel: &checkLevel,
		CheckDepth: &numBlocks,
	}
	return c.sendCmd(cmd)
}

// VerifyChainBlocks requests the server to verify the block chain database
// using the passed check level and number of blocks to verify.
//
// The check level controls how thorough the verification is with higher numbers
// increasing the amount of checks done as consequently how long the
// verification takes.
//
// The number of blocks refers to the number of blocks from the end of the
// current longest chain.
//
// See VerifyChain and VerifyChainLevel to use defaults.
func (c *Client) VerifyChainBlocks(checkLevel, numBlocks int32) (bool, error) {
	return c.VerifyChainBlocksAsync(checkLevel, numBlocks).Receive()
}

// FutureGetTxOutResult is a future promise to deliver the result of a
// GetTxOutAsync RPC invocation (or an applicable error).
type FutureGetTxOutResult chan *response

// Receive waits for the response promised by the future and returns a
// transaction given its hash.
func (r FutureGetTxOutResult) Receive() (*GetTxOutResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// take care of the special case where the output has been spent already
	// it should return the string "null"
	if string(res) == "null" {
		return nil, nil
	}

	// Unmarshal result as an gettxout result object.
	var txOutInfo *GetTxOutResult
	err = json.Unmarshal(res, &txOutInfo)
	if err != nil {
		return nil, err
	}

	return txOutInfo, nil
}

// GetTxOutCmd defines the gettxout JSON-RPC command.
type GetTxOutCmd struct {
	Txid           string
	Vout           uint32
	IncludeMempool *bool `jsonrpcdefault:"true"`
}

// GetTxOutAsync returns an instance of a type that can be used to get
// the result of the RPC at some future time by invoking the Receive function on
// the returned instance.
//
// See GetTxOut for the blocking version and more details.
func (c *Client) GetTxOutAsync(txHash *Hash, index uint32, mempool bool) FutureGetTxOutResult {
	hash := ""
	if txHash != nil {
		hash = txHash.String()
	}

	cmd := &GetTxOutCmd{
		Txid:           hash,
		Vout:           index,
		IncludeMempool: &mempool,
	}
	return c.sendCmd(cmd)
}

// GetTxOut returns the transaction output info if it's unspent and
// nil, otherwise.
func (c *Client) GetTxOut(txHash *Hash, index uint32, mempool bool) (*GetTxOutResult, error) {
	return c.GetTxOutAsync(txHash, index, mempool).Receive()
}

// FutureRescanBlocksResult is a future promise to deliver the result of a
// RescanBlocksAsync RPC invocation (or an applicable error).
//
// NOTE: This is a btcsuite extension ported from
// github.com/decred/dcrrpcclient.
type FutureRescanBlocksResult chan *response

// Receive waits for the response promised by the future and returns the
// discovered rescanblocks data.
//
// NOTE: This is a btcsuite extension ported from
// github.com/decred/dcrrpcclient.
func (r FutureRescanBlocksResult) Receive() ([]RescannedBlock, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	var rescanBlocksResult []RescannedBlock
	err = json.Unmarshal(res, &rescanBlocksResult)
	if err != nil {
		return nil, err
	}

	return rescanBlocksResult, nil
}

// RescanBlocksCmd defines the rescan JSON-RPC command.
//
// NOTE: This is a btcd extension ported from github.com/decred/dcrd/dcrjson
// and requires a websocket connection.
type RescanBlocksCmd struct {
	// Block hashes as a string array.
	BlockHashes []string
}

// RescanBlocksAsync returns an instance of a type that can be used to get the
// result of the RPC at some future time by invoking the Receive function on the
// returned instance.
//
// See RescanBlocks for the blocking version and more details.
//
// NOTE: This is a btcsuite extension ported from
// github.com/decred/dcrrpcclient.
func (c *Client) RescanBlocksAsync(blockHashes []Hash) FutureRescanBlocksResult {
	strBlockHashes := make([]string, len(blockHashes))
	for i := range blockHashes {
		strBlockHashes[i] = blockHashes[i].String()
	}

	cmd := &RescanBlocksCmd{BlockHashes: strBlockHashes}
	return c.sendCmd(cmd)
}

// RescanBlocks rescans the blocks identified by blockHashes, in order, using
// the client's loaded transaction filter.  The blocks do not need to be on the
// main chain, but they do need to be adjacent to each other.
//
// NOTE: This is a btcsuite extension ported from
// github.com/decred/dcrrpcclient.
func (c *Client) RescanBlocks(blockHashes []Hash) ([]RescannedBlock, error) {
	return c.RescanBlocksAsync(blockHashes).Receive()
}

// FutureGetBlockResult is a future promise to deliver the result of a
// GetBlockAsync RPC invocation (or an applicable error).
type FutureGetBlockResult chan *response

// Receive waits for the response promised by the future and returns the raw
// block requested from the server given its hash.
func (r FutureGetBlockResult) Receive() (*MsgBlock, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal result as a string.
	var blockHex string
	err = json.Unmarshal(res, &blockHex)
	if err != nil {
		return nil, err
	}

	// Decode the serialized block hex to raw bytes.
	serializedBlock, err := hex.DecodeString(blockHex)
	if err != nil {
		return nil, err
	}

	// Deserialize the block and return it.
	var msgBlock MsgBlock
	err = msgBlock.Deserialize(bytes.NewReader(serializedBlock))
	if err != nil {
		return nil, err
	}
	return &msgBlock, nil
}

// GetBlockAsync returns an instance of a type that can be used to get the
// result of the RPC at some future time by invoking the Receive function on the
// returned instance.
//
// See GetBlock for the blocking version and more details.
func (c *Client) GetBlockAsync(blockHash *Hash) FutureGetBlockResult {
	hash := ""
	if blockHash != nil {
		hash = blockHash.String()
	}

	cmd := &GetBlockCmd{
		Hash:      hash,
		Verbose:   utils.Basis.Bool(false),
		VerboseTx: nil,
	}
	return c.sendCmd(cmd)
}

// GetBlock returns a raw block from the server given its hash.
//
// See GetBlockVerbose to retrieve a data structure with information about the
// block instead.
func (c *Client) GetBlock(blockHash *Hash) (*MsgBlock, error) {
	return c.GetBlockAsync(blockHash).Receive()
}

// GetBlockVerboseTxAsync returns an instance of a type that can be used to get
// the result of the RPC at some future time by invoking the Receive function on
// the returned instance.
//
// See GetBlockVerboseTx or the blocking version and more details.
func (c *Client) GetBlockVerboseTxAsync(blockHash *Hash) FutureGetBlockVerboseResult {
	hash := ""
	if blockHash != nil {
		hash = blockHash.String()
	}

	cmd := &GetBlockCmd{
		Hash:      hash,
		Verbose:   utils.Basis.Bool(true),
		VerboseTx: utils.Basis.Bool(true),
	}
	return c.sendCmd(cmd)
}

// GetBlockVerboseTx returns a data structure from the server with information
// about a block and its transactions given its hash.
//
// See GetBlockVerbose if only transaction hashes are preferred.
// See GetBlock to retrieve a raw block instead.
func (c *Client) GetBlockVerboseTx(blockHash *Hash) (*GetBlockVerboseResult, error) {
	return c.GetBlockVerboseTxAsync(blockHash).Receive()
}

// FutureGetCFilterResult is a future promise to deliver the result of a
// GetCFilterAsync RPC invocation (or an applicable error).
type FutureGetCFilterResult chan *response

// Receive waits for the response promised by the future and returns the raw
// filter requested from the server given its block hash.
func (r FutureGetCFilterResult) Receive() (*MsgCFilter, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal result as a string.
	var filterHex string
	err = json.Unmarshal(res, &filterHex)
	if err != nil {
		return nil, err
	}

	// Decode the serialized cf hex to raw bytes.
	serializedFilter, err := hex.DecodeString(filterHex)
	if err != nil {
		return nil, err
	}

	// Assign the filter bytes to the correct field of the wire message.
	// We aren't going to set the block hash or extended flag, since we
	// don't actually get that back in the RPC response.
	var msgCFilter MsgCFilter
	msgCFilter.Data = serializedFilter
	return &msgCFilter, nil
}

// FilterType is used to represent a filter type.
type FilterType uint8

// GetCFilterCmd defines the getcfilter JSON-RPC command.
type GetCFilterCmd struct {
	Hash       string
	FilterType FilterType
}

// GetCFilterAsync returns an instance of a type that can be used to get the
// result of the RPC at some future time by invoking the Receive function on the
// returned instance.
//
// See GetCFilter for the blocking version and more details.
func (c *Client) GetCFilterAsync(blockHash *Hash, filterType FilterType) FutureGetCFilterResult {
	hash := ""
	if blockHash != nil {
		hash = blockHash.String()
	}

	cmd := &GetCFilterCmd{
		Hash:       hash,
		FilterType: filterType,
	}
	return c.sendCmd(cmd)
}

//GetCFilter returns a raw filter from the server given its block hash.
func (c *Client) GetCFilter(blockHash *Hash, filterType FilterType) (*MsgCFilter, error) {
	return c.GetCFilterAsync(blockHash, filterType).Receive()
}

//FutureGetCFilterHeaderResult is a future promise to deliver the result of a
//GetCFilterHeaderAsync RPC invocation (or an applicable error).
type FutureGetCFilterHeaderResult chan *response

// Receive waits for the response promised by the future and returns the raw
// filter header requested from the server given its block hash.
func (r FutureGetCFilterHeaderResult) Receive() (*MsgCFHeaders, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal result as a string.
	var headerHex string
	err = json.Unmarshal(res, &headerHex)
	if err != nil {
		return nil, err
	}

	// Assign the decoded header into a hash
	headerHash, err := NewHashFromStr(headerHex)
	if err != nil {
		return nil, err
	}

	// Assign the hash to a headers message and return it.
	msgCFHeaders := MsgCFHeaders{PrevFilterHeader: *headerHash}
	return &msgCFHeaders, nil

}

// GetCFilterHeaderCmd defines the getcfilterheader JSON-RPC command.
type GetCFilterHeaderCmd struct {
	Hash       string
	FilterType FilterType
}

// GetCFilterHeaderAsync returns an instance of a type that can be used to get
// the result of the RPC at some future time by invoking the Receive function
// on the returned instance.
//
// See GetCFilterHeader for the blocking version and more details.
func (c *Client) GetCFilterHeaderAsync(blockHash *Hash,
	filterType FilterType) FutureGetCFilterHeaderResult {
	hash := ""
	if blockHash != nil {
		hash = blockHash.String()
	}

	cmd := &GetCFilterHeaderCmd{
		Hash:       hash,
		FilterType: filterType,
	}
	return c.sendCmd(cmd)
}

// GetCFilterHeader returns a raw filter header from the server given its block
// hash.
func (c *Client) GetCFilterHeader(blockHash *Hash,
	filterType FilterType) (*MsgCFHeaders, error) {
	return c.GetCFilterHeaderAsync(blockHash, filterType).Receive()
}
