defmodule RealDebrid.Helpers do
  @moduledoc """
  Helper functions for working with Real-Debrid data.
  """

  @doc """
  Find a torrent by its hash in a list of torrents.

  Works with maps having either atom or string keys for `:hash` / `"hash"`.

  ## Parameters

    - `torrents` - List of torrent maps
    - `hash` - The torrent hash to find

  ## Returns

    - `torrent` if found
    - `nil` if not found
  """
  @spec get_torrent_by_hash([map()], String.t()) :: map() | nil
  def get_torrent_by_hash(torrents, hash) do
    Enum.find(torrents, fn torrent ->
      Map.get(torrent, :hash) == hash or Map.get(torrent, "hash") == hash
    end)
  end
end
