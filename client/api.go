// Code generated by gen_client. DO NOT EDIT.
package client

import (
	"github.com/pekeps/go-sc2ai/api"
)

func (c *connection) createGame(createGame api.RequestCreateGame) (*api.ResponseCreateGame, error) {
	r, err := c.request(&api.Request{
		Request: &api.Request_CreateGame{
			CreateGame: &createGame,
		},
	})
	return r.GetCreateGame(), err
}

func (c *connection) joinGame(joinGame api.RequestJoinGame) (*api.ResponseJoinGame, error) {
	r, err := c.request(&api.Request{
		Request: &api.Request_JoinGame{
			JoinGame: &joinGame,
		},
	})
	return r.GetJoinGame(), err
}

func (c *connection) restartGame(restartGame api.RequestRestartGame) (*api.ResponseRestartGame, error) {
	r, err := c.request(&api.Request{
		Request: &api.Request_RestartGame{
			RestartGame: &restartGame,
		},
	})
	return r.GetRestartGame(), err
}

func (c *connection) startReplay(startReplay api.RequestStartReplay) (*api.ResponseStartReplay, error) {
	r, err := c.request(&api.Request{
		Request: &api.Request_StartReplay{
			StartReplay: &startReplay,
		},
	})
	return r.GetStartReplay(), err
}

func (c *connection) leaveGame(leaveGame api.RequestLeaveGame) (*api.ResponseLeaveGame, error) {
	r, err := c.request(&api.Request{
		Request: &api.Request_LeaveGame{
			LeaveGame: &leaveGame,
		},
	})
	return r.GetLeaveGame(), err
}

func (c *connection) quickSave(quickSave api.RequestQuickSave) (*api.ResponseQuickSave, error) {
	r, err := c.request(&api.Request{
		Request: &api.Request_QuickSave{
			QuickSave: &quickSave,
		},
	})
	return r.GetQuickSave(), err
}

func (c *connection) quickLoad(quickLoad api.RequestQuickLoad) (*api.ResponseQuickLoad, error) {
	r, err := c.request(&api.Request{
		Request: &api.Request_QuickLoad{
			QuickLoad: &quickLoad,
		},
	})
	return r.GetQuickLoad(), err
}

func (c *connection) quit(quit api.RequestQuit) (*api.ResponseQuit, error) {
	r, err := c.request(&api.Request{
		Request: &api.Request_Quit{
			Quit: &quit,
		},
	})
	return r.GetQuit(), err
}

func (c *connection) gameInfo(gameInfo api.RequestGameInfo) (*api.ResponseGameInfo, error) {
	r, err := c.request(&api.Request{
		Request: &api.Request_GameInfo{
			GameInfo: &gameInfo,
		},
	})
	return r.GetGameInfo(), err
}

func (c *connection) observation(observation api.RequestObservation) (*api.ResponseObservation, error) {
	r, err := c.request(&api.Request{
		Request: &api.Request_Observation{
			Observation: &observation,
		},
	})
	return r.GetObservation(), err
}

func (c *connection) action(action api.RequestAction) (*api.ResponseAction, error) {
	r, err := c.request(&api.Request{
		Request: &api.Request_Action{
			Action: &action,
		},
	})
	return r.GetAction(), err
}

func (c *connection) obsAction(obsAction api.RequestObserverAction) (*api.ResponseObserverAction, error) {
	r, err := c.request(&api.Request{
		Request: &api.Request_ObsAction{
			ObsAction: &obsAction,
		},
	})
	return r.GetObsAction(), err
}

func (c *connection) step(step api.RequestStep) (*api.ResponseStep, error) {
	r, err := c.request(&api.Request{
		Request: &api.Request_Step{
			Step: &step,
		},
	})
	return r.GetStep(), err
}

func (c *connection) data(data api.RequestData) (*api.ResponseData, error) {
	r, err := c.request(&api.Request{
		Request: &api.Request_Data{
			Data: &data,
		},
	})
	return r.GetData(), err
}

func (c *connection) query(query api.RequestQuery) (*api.ResponseQuery, error) {
	r, err := c.request(&api.Request{
		Request: &api.Request_Query{
			Query: &query,
		},
	})
	return r.GetQuery(), err
}

func (c *connection) saveReplay(saveReplay api.RequestSaveReplay) (*api.ResponseSaveReplay, error) {
	r, err := c.request(&api.Request{
		Request: &api.Request_SaveReplay{
			SaveReplay: &saveReplay,
		},
	})
	return r.GetSaveReplay(), err
}

func (c *connection) mapCommand(mapCommand api.RequestMapCommand) (*api.ResponseMapCommand, error) {
	r, err := c.request(&api.Request{
		Request: &api.Request_MapCommand{
			MapCommand: &mapCommand,
		},
	})
	return r.GetMapCommand(), err
}

func (c *connection) replayInfo(replayInfo api.RequestReplayInfo) (*api.ResponseReplayInfo, error) {
	r, err := c.request(&api.Request{
		Request: &api.Request_ReplayInfo{
			ReplayInfo: &replayInfo,
		},
	})
	return r.GetReplayInfo(), err
}

func (c *connection) availableMaps(availableMaps api.RequestAvailableMaps) (*api.ResponseAvailableMaps, error) {
	r, err := c.request(&api.Request{
		Request: &api.Request_AvailableMaps{
			AvailableMaps: &availableMaps,
		},
	})
	return r.GetAvailableMaps(), err
}

func (c *connection) saveMap(saveMap api.RequestSaveMap) (*api.ResponseSaveMap, error) {
	r, err := c.request(&api.Request{
		Request: &api.Request_SaveMap{
			SaveMap: &saveMap,
		},
	})
	return r.GetSaveMap(), err
}

func (c *connection) ping(ping api.RequestPing) (*api.ResponsePing, error) {
	r, err := c.request(&api.Request{
		Request: &api.Request_Ping{
			Ping: &ping,
		},
	})
	return r.GetPing(), err
}

func (c *connection) debug(debug api.RequestDebug) (*api.ResponseDebug, error) {
	r, err := c.request(&api.Request{
		Request: &api.Request_Debug{
			Debug: &debug,
		},
	})
	return r.GetDebug(), err
}
