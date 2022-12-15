import { useAppDispatch, useAppSelector } from "../../hooks/reduxHooks";
import { SideNavProps } from "../../types/dashboard";
import { setPage } from "../reducer";

export default function SideNav({ groups }: SideNavProps) {
  const state = useAppSelector((state) => state.dashboard);
  const dispatch = useAppDispatch();

  return (
    <div className="Dashboard-side-nav">
      {groups.map((group) => (
        <div className="Dashboard-side-nav-group" key={group.groupName}>
          {state.isSideNavOpen && <h1>{group.groupName}</h1>}
          {group.options.map((option) => (
            <div
              className="Side-nav-button"
              key={option.name}
              onClick={() => dispatch(setPage(option.page))}
              id={
                state.page === option.page
                  ? "SideNavButtonSelected"
                  : "SideNavButtonUnselected"
              }
            >
              <i>{option.icon}</i>
              {state.isSideNavOpen && <p>{option.name}</p>}
            </div>
          ))}
          {!state.isSideNavOpen && <hr />}
        </div>
      ))}
    </div>
  );
}
